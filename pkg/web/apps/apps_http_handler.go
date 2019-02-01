package apps

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// HTTPHandler instance
	HTTPHandler struct {
		Service Service
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, s Service) *HTTPHandler {
	handler := &HTTPHandler{
		Service: s,
	}

	return handler
}

// GetApps returns all apps as JSON from the AppService
func (a *HTTPHandler) GetApps(c echo.Context) error {
	apps, err := a.Service.GetApps(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apps)
}
