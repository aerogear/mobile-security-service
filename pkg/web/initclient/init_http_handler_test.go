package initclient

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/helpers"

	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/labstack/echo"
)

func TestHTTPHandler_InitClientApp(t *testing.T) {
	validDevice := helpers.GetMockDevice()
	validDevice.Version = "1.0.0"

	// set up mock models
	deviceWithoutDeviceID := *validDevice
	deviceWithoutDeviceID.DeviceID = ""

	deviceWithoutVersion := *validDevice
	deviceWithoutVersion.Version = ""

	deviceWithoutAppID := *validDevice
	deviceWithoutAppID.AppID = ""

	type fields struct {
		appsService apps.Service
	}
	type args struct {
		device models.Device
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		mockAppService *apps.ServiceMock
	}{
		{
			name: "A 400 Bad Request should be returned when request body is missing device ID",
			args: args{
				device: deviceWithoutDeviceID,
			},
			mockAppService: &apps.ServiceMock{
				InitClientAppFunc: func(device *models.Device) (*models.Version, error) {
					return nil, nil
				},
			},
			wantStatusCode: 400,
		},
		{
			name: "A 400 Bad Request should be returned when request body is missing version",
			args: args{
				device: deviceWithoutVersion,
			},
			mockAppService: &apps.ServiceMock{
				InitClientAppFunc: func(device *models.Device) (*models.Version, error) {
					return nil, nil
				},
			},
			wantStatusCode: 400,
		},
		{
			name: "A 400 Bad Request should be returned when request body is missing app ID",
			args: args{
				device: deviceWithoutAppID,
			},
			wantStatusCode: 400,
			mockAppService: &apps.ServiceMock{
				InitClientAppFunc: func(device *models.Device) (*models.Version, error) {
					return nil, nil
				},
			},
		},
		{
			name: "Expect init data to be returned when valid device is supplied",
			args: args{
				device: *validDevice,
			},
			mockAppService: &apps.ServiceMock{
				InitClientAppFunc: func(device *models.Device) (*models.Version, error) {
					return &models.Version{
						ID:              uuid.New().String(),
						Version:         device.Version,
						AppID:           device.AppID,
						Disabled:        true,
						DisabledMessage: "App is disabled",
					}, nil
				},
			},
			wantStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			deviceJSON, _ := json.Marshal(tt.args.device)

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(deviceJSON)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/init")

			handler := NewHTTPHandler(e, tt.mockAppService)

			if handler.InitClientApp(c); rec.Code != tt.wantStatusCode {
				t.Errorf("HTTPHandler.InitClientApp() statusCode = %v, wantStatusCode %v", rec.Code, tt.wantStatusCode)
			}
		})
	}
}

func trimBody(body string) string {
	return strings.TrimSpace(body)
}
