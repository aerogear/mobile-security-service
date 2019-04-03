package apps

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

var (
	// make and configure a mocked Service which will return the success scenarios
	mockedService = &ServiceMock{
		DisableAllAppVersionsByAppIDFunc: func(id string, message string) error {
			return nil
		},
		GetActiveAppByIDFunc: func(ID string) (*models.App, error) {
			app := helpers.GetMockApp()
			if app.ID == ID {
				return app, nil
			}
			return nil, models.ErrNotFound
		},
		GetActiveAppByAppIDFunc: func(appId string) (*models.App, error) {
			app := helpers.GetMockApp()
			if app.AppID == appId {
				return app, nil
			}
			return nil, models.ErrNotFound
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				*helpers.GetMockApp(),
			}, nil
		},
		UpdateAppVersionsFunc: func(id string, versions []models.Version) error {
			return nil
		},
		DeleteAppByIdFunc: func(id string) error {
			return nil
		},
		CreateAppFunc: func(app models.App) error {
			return nil
		},
		UpdateAppNameByIDFunc: func(id string, name string) error {
			return nil
		},
	}

	// make and configure a mocked Service which will return the scenarios with errors
	mockedServiceWithError = &ServiceMock{
		DisableAllAppVersionsByAppIDFunc: func(id string, message string) error {
			return models.ErrInternalServerError
		},
		GetActiveAppByIDFunc: func(ID string) (*models.App, error) {
			return nil, models.ErrInternalServerError
		},
		GetActiveAppByAppIDFunc: func(appId string) (*models.App, error) {
			return nil, models.ErrNotFound
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrNotFound
		},
		UpdateAppVersionsFunc: func(id string, versions []models.Version) error {
			return models.ErrNotFound
		},
		DeleteAppByIdFunc: func(id string) error {
			return models.ErrInternalServerError
		},
		CreateAppFunc: func(app models.App) error {
			return models.ErrInternalServerError
		},
		UpdateAppNameByIDFunc: func(id string, name string) error {
			return models.ErrNotFound
		},
	}
)

func Test_HttpHandler_GetApps(t *testing.T) {
	// make and configure a mocked Service which will return the scenarios with errors
	mockedServiceWithInternalError := &ServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrInternalServerError
		},
	}
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		wantErr     bool
		wantCode    int
		mockService ServiceMock
	}{
		{
			name:        "Should return success to get all",
			wantErr:     false,
			mockService: *mockedService,
			wantCode:    200,
		},
		{
			name:        "Should return when no apps have been found, return a HTTP Status code of 204 with no response body",
			wantErr:     true,
			mockService: *mockedServiceWithError,
			wantCode:    204,
		},
		{
			name:        "Should return error when an error occurs in the database",
			wantErr:     true,
			mockService: *mockedServiceWithInternalError,
			wantCode:    500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/apps")
			h := NewHTTPHandler(e, &tt.mockService)
			err := h.GetApps(c)
			if err != nil {
				t.Errorf("httpHandler.GetApps() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.GetApps() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_GetAppsWithQueryParameter(t *testing.T) {
	mockedServiceWithInternalError := &ServiceMock{
		GetActiveAppByAppIDFunc: func(appId string) (*models.App, error) {
			return nil, models.ErrInternalServerError
		},
	}

	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		appId       string
		wantErr     bool
		wantCode    int
		mockService ServiceMock
	}{
		{
			name:        "Should return success to get the app by appId",
			wantErr:     false,
			appId:       helpers.GetMockApp().AppID,
			mockService: *mockedService,
			wantCode:    200,
		},
		{
			name:        "Should return when no app by appId have been found, return a HTTP Status code of 204 with no response body",
			wantErr:     true,
			appId:       helpers.GetMockApp().AppID,
			mockService: *mockedServiceWithError,
			wantCode:    204,
		},
		{
			name:        "Should return error when an error occurs in the database",
			wantErr:     true,
			appId:       helpers.GetMockApp().AppID,
			mockService: *mockedServiceWithInternalError,
			wantCode:    500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			q := make(url.Values)
			q.Set("appId", tt.appId)
			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/apps")
			h := NewHTTPHandler(e, &tt.mockService)
			err := h.GetApps(c)
			if err != nil {
				t.Errorf("httpHandler.GetApps() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.GetApps() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_DeleteAppById(t *testing.T) {
	// make and configure a mocked Service which will return the scenarios with errors
	mockedServiceWithErroNotFound := &ServiceMock{
		DeleteAppByIdFunc: func(id string) error {
			return models.ErrNotFound
		},
	}
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		id          string
		fields      fields
		wantErr     bool
		wantCode    int
		mockService ServiceMock
	}{
		{
			name:        "Should return success to delete app by id",
			wantErr:     false,
			id:          helpers.GetMockApp().ID,
			mockService: *mockedService,
			wantCode:    204,
		},
		{
			name:        "Should return error when no apps have been found",
			wantErr:     true,
			id:          helpers.GetMockApp().ID,
			mockService: *mockedServiceWithErroNotFound,
			wantCode:    404,
		},
		{
			name:        "Should return error when an error occurs in the database",
			wantErr:     true,
			id:          helpers.GetMockApp().ID,
			mockService: *mockedServiceWithError,
			wantCode:    500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/apps")
			c.SetParamNames("id")
			c.SetParamValues(tt.id)
			h := NewHTTPHandler(e, &tt.mockService)
			err := h.DeleteAppById(c)
			if err != nil {
				t.Errorf("httpHandler.DeleteAppById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.DeleteAppById() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_httpHandler_UpdateAllAppVersionsByAppID(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		id          string
		wantErr     bool
		wantCode    int
		data        []models.Version
		mockService ServiceMock
	}{
		{
			name:        "Should update the versions with success",
			id:          helpers.GetMockApp().ID,
			data:        helpers.GetMockAppVersionList(),
			wantErr:     false,
			mockService: *mockedService,
			wantCode:    204,
		},
		{
			name:        "Should return error since it is an invalid id",
			id:          "invalid",
			data:        helpers.GetMockAppVersionList(),
			wantErr:     true,
			mockService: *mockedService,
			wantCode:    400,
		},
		{
			name:        "Update all versions should return return error with an empty collection of data sent",
			id:          helpers.GetMockApp().ID,
			data:        []models.Version{},
			wantErr:     true,
			mockService: *mockedService,
			wantCode:    400,
		},
		{
			name:        "Update all versions should return return error in the database",
			id:          helpers.GetMockApp().ID,
			data:        helpers.GetMockAppVersionList(),
			wantErr:     true,
			mockService: *mockedServiceWithError,
			wantCode:    404,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		allVersions, _ := json.Marshal(tt.data)
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(allVersions)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id/versions")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, &tt.mockService)
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			err := h.UpdateAppVersions(c)
			if err != nil {
				t.Errorf("httpHandler.UpdateAppVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.UpdateAppVersions() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_httpHandler_UpdateAppNameByID(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		id          string
		wantErr     bool
		wantCode    int
		data        models.App
		mockService ServiceMock
	}{
		{
			name:        "Should update the name with success",
			id:          helpers.GetMockApp().ID,
			data:        *helpers.GetMockApp(),
			wantErr:     false,
			mockService: *mockedService,
			wantCode:    204,
		},
		{
			name:        "Should return error since it is an invalid id",
			id:          "invalid",
			data:        *helpers.GetMockApp(),
			wantErr:     true,
			mockService: *mockedService,
			wantCode:    400,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		app, _ := json.Marshal(tt.data)
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(string(app)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, &tt.mockService)
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			err := h.UpdateAppNameByID(c)
			if err != nil {
				t.Errorf("httpHandler.UpdateAppNameByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.UpdateAppNameByID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_UpdateAllAppVersionsByAppID_WithInvalidJsonData(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name     string
		fields   fields
		id       string
		wantErr  bool
		want     int
		wantCode int
	}{
		{
			name:     "Update all versions should return return error with invalid data sent",
			id:       helpers.GetMockApp().ID,
			want:     http.StatusBadRequest,
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string("")))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id/versions")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, mockedService)
		c.SetParamValues(tt.id)

		t.Run(tt.name, func(t *testing.T) {
			h.UpdateAppVersions(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.UpdateAppVersions() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_DisableAllAppVersionsByAppID(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		id          string
		wantErr     bool
		wantCode    int
		data        models.Version
		mockService ServiceMock
	}{
		{
			name:        "Disable all versions should return success",
			id:          helpers.GetMockApp().ID,
			data:        helpers.GetMockAppVersionForDisableAll(),
			mockService: *mockedService,
			wantErr:     false,
			wantCode:    200,
		},
		{
			name:        "Disable all versions should return error with invalid ID",
			id:          "invalid",
			data:        helpers.GetMockAppVersionForDisableAll(),
			mockService: *mockedService,
			wantErr:     true,
			wantCode:    400,
		},
		{
			name:        "Disable all versions should return error from the database",
			id:          helpers.GetMockApp().ID,
			mockService: *mockedServiceWithError,
			wantErr:     true,
			wantCode:    500,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		version, _ := json.Marshal(tt.data)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(version)))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id/versions/disable")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, &tt.mockService)
		c.SetParamValues(tt.id)

		t.Run(tt.name, func(t *testing.T) {
			err := h.DisableAllAppVersionsByAppID(c)
			if err != nil {
				t.Errorf("httpHandler.DisableAllAppVersionsByAppID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.DisableAllAppVersionsByAppID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_DisableAllAppVersionsByAppID_WithInvalidJsonData(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name     string
		fields   fields
		id       string
		wantErr  bool
		want     int
		wantCode int
	}{
		{
			name:     "Disable all versions should return return error with invalid data sent",
			id:       helpers.GetMockApp().ID,
			want:     http.StatusBadRequest,
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string("")))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps/:id/versions/disable")
		c.SetParamNames("id")
		h := NewHTTPHandler(e, mockedService)
		c.SetParamValues(tt.id)

		t.Run(tt.name, func(t *testing.T) {
			h.DisableAllAppVersionsByAppID(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.DisableAllAppVersionsByAppID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_httpHandler_GetActiveAppByID(t *testing.T) {
	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	// make and configure a mocked Service
	app := helpers.GetMockApp()

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
			wantCode: 200,
		},
		{
			name:     "Get app by id using an invalid id format should return an error",
			id:       "some string that should fail",
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
		h := NewHTTPHandler(e, mockedService)
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			_ = h.GetActiveAppByID(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.GetActiveAppByID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_CreateApp(t *testing.T) {
	mockAppWithoutName := &models.App{
		AppID: "com.aerogear.mobile_app_one",
	}

	mockAppWithoutID := &models.App{
		AppName: "Mobile App One",
	}

	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		data        *models.App
		appId       string
		wantErr     bool
		want        int
		wantCode    int
		mockService ServiceMock
	}{
		{
			name:        "Create with appId and name should return success",
			data:        helpers.GetMockApp(),
			want:        http.StatusNoContent,
			mockService: *mockedService,
			wantCode:    201,
		},
		{
			name:        "Create with just appId should return success",
			data:        mockAppWithoutName,
			want:        http.StatusNoContent,
			mockService: *mockedService,
			wantCode:    201,
		},
		{
			name:        "Create app without appId should return error",
			data:        mockAppWithoutID,
			want:        http.StatusBadRequest,
			mockService: *mockedService,
			wantCode:    400,
		},
		{
			name:        "Error in the database should return error in the post request",
			data:        helpers.GetMockApp(),
			want:        http.StatusInternalServerError,
			mockService: *mockedServiceWithError,
			wantCode:    500,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		app, _ := json.Marshal(tt.data)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(app)))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/apps")
		h := NewHTTPHandler(e, &tt.mockService)

		t.Run(tt.name, func(t *testing.T) {
			h.CreateApp(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.CreateApp() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}
