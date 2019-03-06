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
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				*helpers.GetMockApp(),
			}, nil
		},
		UpdateAppVersionsFunc: func(versions []models.Version) error {
			return nil
		},
		InitClientAppFunc: func(sdkInfo *models.Device) (*models.Version, error) {
			return &helpers.GetMockAppVersionList()[0], nil // FIXME: Create an only object result
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
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrNotFound
		},
		UpdateAppVersionsFunc: func(versions []models.Version) error {
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
			wantCode:    200,
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
		c.SetPath(APIRoutePrefix + "/apps/:id/versions/batch-disable")
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
		c.SetPath(APIRoutePrefix + "/apps/:id/versions/batch-disable")
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
		h := NewHTTPHandler(e, mockedService)
		c.SetParamValues(tt.id)
		t.Run(tt.name, func(t *testing.T) {
			_ = h.GetActiveAppByID(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.GetActiveAppByID() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
			if strings.TrimSpace(rec.Body.String()) != tt.want {
				t.Errorf("httpHandler.GetActiveAppByID() got %v, want %v", strings.TrimSpace(rec.Body.String()), tt.want)
			}
		})
	}
}

func TestHTTPHandler_InitClientApp(t *testing.T) {
	deviceWithMissingDeviceID := helpers.GetMockDevice() // See that just one mock is enough
	deviceWithMissingDeviceID.DeviceID = ""              // See here that you can change the object to match with the criteria

	deviceWithMissingVersion := helpers.GetMockDevice()
	deviceWithMissingVersion.VersionID = ""

	config := config.Get()
	APIRoutePrefix := config.APIRoutePrefix
	type fields struct {
		Service Service
	}
	tests := []struct {
		name        string
		fields      fields
		device      models.Device
		mockService ServiceMock
		wantErr     error
		wantCode    int
	}{
		{
			name:        "A 400 Bad Request should be returned when request body is missing device ID",
			device:      *deviceWithMissingDeviceID,
			mockService: *mockedService,
			wantCode:    400,
		},
		{
			name:        "A 400 Bad Request should be returned when request body is missing version",
			device:      *deviceWithMissingVersion,
			mockService: *mockedService,
			wantCode:    400,
		},
	}
	for _, tt := range tests {
		device, _ := json.Marshal(tt.device)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(device))) // body -> device
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(APIRoutePrefix + "/init")
		h := NewHTTPHandler(e, mockedService)
		t.Run(tt.name, func(t *testing.T) {
			h.InitClientApp(c)
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.InitClientApp() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}
