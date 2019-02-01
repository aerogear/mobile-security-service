package apps

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

var mockPostgresRepository = NewPostgreSQLRepository()
var mockService = NewService(mockPostgresRepository)

func TestHTTPHandler_GetApps(t *testing.T) {
	// set up mock context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/apps")
	h := NewHTTPHandler(e, mockService)

	_ = h.GetApps(c)

	if rec.Code != http.StatusOK {
		t.Errorf("HTTPHandler.GetApps() statusCode = %v, wantCode = %v", rec.Code, http.StatusOK)
	}

	expected := `[{"id":1,"appId":"com.aerogear.app1","appName":"app1","numOfDeployedVersions":1,"numOfClients":1,"numOfAppLaunches":1}]`

	resBody := trimBody(rec.Body.String())

	if resBody != expected {
		t.Errorf("HTTPHandler.GetApps() want = %v, wantCode = %v", expected, resBody)
	}
}

func trimBody(body string) string {
	return strings.TrimSpace(body)
}
