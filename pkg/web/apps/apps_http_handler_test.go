package apps

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_HttpHandler_GetApps(t *testing.T) {
	// make and configure a mocked Service
	mockedAppService := &ServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				*helpers.GetMockApp(),
			}, nil
		},
	}

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/apps")
	h := &httpHandler{mockedAppService}
	// Assertions
	if assert.NoError(t, h.GetApps(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

//TODO when update app endpoint created its hitting get apps by id endpoint at the moment
func Test_HttpHandler_UpdateApp(t *testing.T) {
	// make and configure a mocked Service
	app := helpers.GetMockApp()
	mockedAppService := &ServiceMock{
		GetAppByIDFunc: func(ID string) (*models.App, error) {
			return app, nil
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				*helpers.GetMockApp(),
			}, nil
		},
	}
	// Setup
	e := echo.New()
	appJson, _ := json.Marshal(helpers.GetMockApp())
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(appJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/apps/:id")
	c.SetParamNames("id")
	c.SetParamValues(app.ID)
	h := &httpHandler{mockedAppService}

	// Assertions
	if assert.NoError(t, h.UpdateApp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: It will be fixed when the handler be implemented
		//assert.Equal(t, string(appJson)+"\n", rec.Body.String())
	}
}

func Test_httpHandler_GetAppByID(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	// make and configure a mocked Service
	app := helpers.GetMockApp()
	mockedAppService := &ServiceMock{
		GetAppByIDFunc: func(ID string) (*models.App, error) {
			if app.ID == ID {
				return app, nil
			}
			return nil, models.ErrNotFound
		},
	}
	// Setup

	type fields struct {
		Service Service
	}
	tests := []struct {
		name     string
		fields   fields
		id       string
		wantErr  bool
		wantCode int
		want     string
	}{
		{
			name:     "Get app by id valid format but id not found",
			id:       "cb1becdd-7726-4902-9014-6fb2296b9ae6",
			want:     `{"message":"Your requested Item is not found","statusCode":404}`,
			wantCode: 404,
		},
		{
			name:     "Get app by id should return an app",
			id:       app.ID,
			want:     `{"id":"7f89ce49-a736-459e-9110-e52d049fc025","appId":"com.aerogear.mobile_app_one","appName":"Mobile App One","deployedVersions":[{"id":"55ebd387-9c68-4137-a367-a12025cc2cdb","version":"1.0","appId":"com.aerogear.mobile_app_one","disabled":false,"disabledMessage":"Please contact an administrator","numOfCurrentInstalls":1,"numOfAppLaunches":2},{"id":"59ebd387-9c68-4137-a367-a12025cc1cdb","version":"1.1","appId":"com.aerogear.mobile_app_one","disabled":false,"numOfCurrentInstalls":0,"numOfAppLaunches":0},{"id":"59dbd387-9c68-4137-a367-a12025cc2cdb","version":"1.0","appId":"com.aerogear.mobile_app_two","disabled":false,"numOfCurrentInstalls":0,"numOfAppLaunches":0}]}`,
			wantCode: 200,
		},
		{
			name:     "Get app by id using an invalid id format should return an error",
			id:       "some string that should fail",
			want:     `{"message":"Invalid id supplied","statusCode":400}`,
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, mockedAppService)
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			_ = h.GetAppByID(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.GetAppByID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
			if strings.TrimSpace(rec.Body.String()) != tt.want {
				t.Errorf("httpHandler.GetAppByID() got %v, want %v", strings.TrimSpace(rec.Body.String()), tt.want)
			}
		})
	}
}
