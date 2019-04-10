package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func Test_HttpHandler_GetUser(t *testing.T) {

	type headers struct {
		username string
		email    string
	}
	tests := []struct {
		name     string
		wantErr  bool
		wantCode int
		headers  headers
	}{
		{
			name:     "Should return success when header provided",
			wantErr:  false,
			wantCode: 200,
			headers: headers{
				username: "TestUser",
				email:    "test@user.com",
			},
		},
		{
			name:     "Should return notfound, when no headers are provided",
			wantErr:  true,
			wantCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if !tt.wantErr {
				req.Header.Add(USER_NAME_HEADER, tt.headers.username)
				req.Header.Add(USER_EMAIL_HEADER, tt.headers.email)
			}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/user")
			h := NewHTTPHandler(e)
			err := h.GetUser(c)
			if err != nil {
				t.Errorf("httpHandler.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantCode {
				t.Errorf("httpHandler.Get User() statusCode = %v, wantcode = %v", rec.Code, tt.wantCode)
			}
		})
	}
}
