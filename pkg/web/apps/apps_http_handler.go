package apps

import (
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/aerogear/mobile-security-service/pkg/httperrors"
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
	apps, err := a.Service.GetApps()

	// If no apps have been found, return a HTTP Status code of 204 with no response body
	if err == models.ErrNotFound {
		return c.NoContent(http.StatusNoContent)
	}

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, apps)
}
