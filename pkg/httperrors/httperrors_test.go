package httperrors

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

func TestHTTPError(t *testing.T) {

	type args struct {
		statusCode int
		message    string
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantMessage string
	}{
		{
			name:     "HTTPError() should return a HTTP error with the provided statusCode arg",
			args:     args{400, ""},
			wantCode: 400,
		},
		{
			name:        "HTTPErrors() should return a 500 Internal Server when an invalid HTTP Status Code was supplied",
			args:        args{100, ""},
			wantCode:    500,
			wantMessage: "Invalid HTTP status code",
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
				t.Errorf("HTTPError() statusCode = %v, wantCode %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("HTTPError() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.wantMessage == "" {
				tt.wantMessage = tt.args.message
			}

			if tt.wantMessage != responseBody.Message {
				t.Errorf("HTTPError() wantMessage = %v, got = %v", tt.wantMessage, responseBody.Message)
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
				t.Errorf("GetHTTPResponseFromErr() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("GetHTTPResponseFromErr() could not unmarshal response body into errResponse struct")
			}
		})
	}
}

func TestBadRequest(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "BadRequest() should return a 400 response code with the default error message",
			args:     args{""},
			wantCode: 400,
		},
		{
			name:     "BadRequest() should return a 400 response code with a custom error message",
			args:     args{"Bad request made to the server"},
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = BadRequest(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("BadRequest() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("BadRequest() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("BadRequest() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestUnauthorized(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "Unauthorized() should return a 401 response code with the default error message",
			args:     args{""},
			wantCode: 401,
		},
		{
			name:     "Unauthorized() should return a 401 response code with a custom error message",
			args:     args{"You are unauthorized to view this resource"},
			wantCode: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = Unauthorized(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("HTTPError() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("Unauthorized() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("Unauthorized() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestForbidden(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "Forbidden() should return a 403 response code with the default error message",
			args:     args{""},
			wantCode: 403,
		},
		{
			name:     "Forbidden() should return a 403 response code with a custom error message",
			args:     args{"You are forbidden to view this resource"},
			wantCode: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = Forbidden(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("Forbidden() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("Forbidden() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("Forbidden() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "NotFound() should return a 404 response code with the default error message",
			args:     args{""},
			wantCode: 404,
		},
		{
			name:     "NotFound() should return a 404 response code with a custom error message",
			args:     args{"Resource not found"},
			wantCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = NotFound(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("NotFound() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("NotFound() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("NotFound() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "MethodNotAllowed() should return a 405 response code with the default error message",
			args:     args{""},
			wantCode: 405,
		},
		{
			name:     "MethodNotAllowed() should return a 405 response code with a custom error message",
			args:     args{"This method is not allowed"},
			wantCode: 405,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = MethodNotAllowed(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("MethodNotAllowed() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("MethodNotAllowed() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("MethodNotAllowed() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestConflict(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "Conflict() should return a 409 response code with the default error message",
			args:     args{""},
			wantCode: 409,
		},
		{
			name:     "MethodNotAllowed() should return a 409 response code with a custom error message",
			args:     args{"Conflict Found"},
			wantCode: 409,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = Conflict(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("Conflict() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("Conflict() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("Conflict() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestGone(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "Gone() should return a 410 response code with the default error message",
			args:     args{""},
			wantCode: 410,
		},
		{
			name:     "MethodNotAllowed() should return a 410 response code with a custom error message",
			args:     args{"Resource is gone"},
			wantCode: 410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = Gone(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("Gone() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("Gone() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("Gone() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestUnsupportedMediaType(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "UnsupportedMediaType() should return a 409 response code with the default error message",
			args:     args{""},
			wantCode: 415,
		},
		{
			name:     "UnsupportedMediaType() should return a 409 response code with a custom error message",
			args:     args{"This media type is unsupported"},
			wantCode: 415,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = UnsupportedMediaType(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("UnsupportedMediaType() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("UnsupportedMediaType() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("UnsupportedMediaType() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestInternalServerError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "InternalServerError() should return a 500 response code with the default error message",
			args:     args{""},
			wantCode: 500,
		},
		{
			name:     "InternalServerError() should return a 500 response code with a custom error message",
			args:     args{"An error has occurred in the server"},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = InternalServerError(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("InternalServerError() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("InternalServerError() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("InternalServerError() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}

func TestNotImplemented(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "NotImplemented() should return a 501 response code with the default error message",
			args:     args{""},
			wantCode: 501,
		},
		{
			name:     "NotImplemented() should return a 501 response code with a custom error message",
			args:     args{"An error has occurred in the server"},
			wantCode: 501,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if _ = NotImplemented(c, tt.args.message); !reflect.DeepEqual(rec.Code, tt.wantCode) {
				t.Errorf("NotImplemented() error = %v, wantErr %v", rec.Code, tt.wantCode)
			}

			// Unmarshal the raw response body into errResponse struct
			responseBody := errResponse{}
			b := []byte(rec.Body.String())

			if err := json.Unmarshal(b, &responseBody); err != nil {
				t.Errorf("NotImplemented() could not unmarshal response body into errResponse struct")
			}

			// if the message arg is empty, use the default for this status code
			if tt.args.message == "" {
				tt.args.message = codes[tt.wantCode]
			}

			if tt.args.message != responseBody.Message {
				t.Errorf("NotImplemented() wantMessage = %v, got = %v", tt.args.message, responseBody.Message)
			}
		})
	}
}
