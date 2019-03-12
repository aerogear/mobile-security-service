package checks

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// make and configure a mocked Service which will return the success scenarios
	mockedServiceSuccess = &apps.ServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				*helpers.GetMockApp(),
			}, nil
		},
	}

	mockedServiceNoData = &apps.ServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrNotFound
		},
	}

	mockedServiceError = &apps.ServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrInternalServerError
		},
	}
)

func Test_HttpHandler_Ping(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		wantCode int
	}{
		{
			name:     "Should return success",
			wantErr:  false,
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/ping")
			h := NewHTTPHandler(e, mockedServiceSuccess)
			err := h.Ping(c)
			if err != nil {
				t.Errorf("httpHandler.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.Ping() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}

func Test_HttpHandler_Healthz(t *testing.T) {
	type fields struct {
		Service apps.Service
	}
	tests := []struct {
		name        string
		fields      fields
		wantErr     bool
		wantCode    int
		mockService apps.ServiceMock
	}{
		{
			name:        "Should return success",
			wantErr:     false,
			mockService: *mockedServiceSuccess,
			wantCode:    200,
		},
		{
			name:        "Should return success",
			wantErr:     false,
			mockService: *mockedServiceNoData,
			wantCode:    200,
		},
		{
			name:        "Should return error",
			wantErr:     true,
			mockService: *mockedServiceError,
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
			c.SetPath("/api/healthz")
			h := NewHTTPHandler(e, &tt.mockService)
			err := h.Healthz(c)
			if err != nil {
				t.Errorf("httpHandler.Health() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("HTTPHandler.Health() statusCode = %v, wantCode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}
