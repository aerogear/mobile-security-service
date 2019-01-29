package httperrors

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

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
