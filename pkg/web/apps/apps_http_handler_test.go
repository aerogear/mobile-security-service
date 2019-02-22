package apps

import (
	"encoding/json"
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HttpHandler_GetApps(t *testing.T) {
	// make and configure a mocked Service
	mockedAppService := &AppServiceMock{
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

func Test_HttpHandler_GetAppsById(t *testing.T) {
	// make and configure a mocked Service
	mockedAppService := &AppServiceMock{
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
	c.SetPath("/api/apps/:id")
	c.SetParamNames("id")
	c.SetParamValues(helpers.GetMockApp().AppID)
	h := &httpHandler{mockedAppService}

	// Assertions
	if assert.NoError(t, h.GetAppByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: It will be fixed when the handler be implemented
		// appJson, _ := json.Marshal(helpers.GetMockApp())
		//assert.Equal(t, string(appJson)+"\n", rec.Body.String())
	}
}

func Test_HttpHandler_UpdateApp(t *testing.T) {
	// make and configure a mocked Service
	mockedAppService := &AppServiceMock{
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
	h := &httpHandler{mockedAppService}

	// Assertions
	if assert.NoError(t, h.UpdateApp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: It will be fixed when the handler be implemented
		//assert.Equal(t, string(appJson)+"\n", rec.Body.String())
	}
}
