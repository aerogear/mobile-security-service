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

func Test_HttpHandler_GetAppsByID_mock(t *testing.T) {
	config := config.Get()
  APIRoutePrefix := config.APIRoutePrefix


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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(APIRoutePrefix + "/apps/:id")
	c.SetParamNames("id")
	c.SetParamValues(app.ID)
	h := NewHTTPHandler(e, mockedAppService)

	// Assertions
	if assert.NoError(t, h.GetAppByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		appJSON, _ := json.Marshal(helpers.GetMockApp())
		assert.Equal(t, string(appJSON)+"\n", rec.Body.String())
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
			if app != nil { 
				return app, nil
			}
			return nil, models.ErrNotFound
		},
	}
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(APIRoutePrefix + "/apps/:id")
	c.SetParamNames("id")
	h := NewHTTPHandler(e, mockedAppService)
	type fields struct {
		Service Service
	}
	tests := []struct {
		name    string
		fields  fields
		id      string
		wantErr bool
		want    *models.App
	}{
		{
			name: "Get app by id should return an app",
			id: app.ID,
			want: app,
		},
		{
			name: "Get app by id using an invalid id format should return an error",
			id: "some string that should fail",
			want: nil,
		},
		{
			name:"Get app by id valid format but id not found",
			id: "cb1becdd-7726-4902-9014-6fb2296b9ae6",
			want: nil,
		},
	}
	for _, tt := range tests {
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			if err := h.GetAppByID(c); (err != nil) != tt.wantErr {
				t.Errorf("httpHandler.GetAppByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
