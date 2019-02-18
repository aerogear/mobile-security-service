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
		Service AppService
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, s AppService) *HTTPHandler {
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

// TODO: Implement
//GetAppById returns a app with the ID in JSON format from the AppService
func (a *HTTPHandler) GetAppById(c echo.Context) error {
	apps, err := a.Service.GetApps()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apps)
}

// TODO: Implement
//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *HTTPHandler) UpdateApp(c echo.Context) error {
	apps, err := a.Service.GetApps()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apps)
}
