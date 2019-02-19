package httperrors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/models"
	"github.com/labstack/echo"
)

func TestHTTPError(t *testing.T) {

	type args struct {
		statusCode int
		message    string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "HTTPError() should return a HTTP error with the provided statusCode arg",
			args:     args{400, ""},
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = HTTPError(c, tt.args.statusCode, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("HTTPError() error = %v, wantErr %v", rec.Code, tt.args.statusCode)
			}
		})
	}
}

func TestGetHTTPResponseFromErr(t *testing.T) {

	tests := []struct {
		name     string
		args     error
		wantCode int
	}{
		{
			name:     "GetHTTPResponseFromErr() should return a HTTP error with a 500 status code",
			args:     models.ErrInternalServerError,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "GetHTTPResponseFromErr() should return a HTTP error with a 404 status code",
			args:     models.ErrNotFound,
			wantCode: http.StatusNotFound,
		},
		{
			name:     "GetHTTPResponseFromErr() should return a HTTP error with a 409 status code",
			args:     models.ErrConflict,
			wantCode: http.StatusConflict,
		},
		{
			name:     "GetHTTPResponseFromErr() should return a HTTP error with a 400 status code",
			args:     models.ErrBadParamInput,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "GetHTTPResponseFromErr() should return a HTTP error with a 401 status code",
			args:     models.ErrUnauthorized,
			wantCode: http.StatusUnauthorized,
		},
		{
			name:     "GetHTTPResponseFromErr() should return a default HTTP error with a 500 status code",
			args:     errors.New(""),
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = GetHTTPResponseFromErr(c, tt.args); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("HTTPError() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}
		})
	}
}
