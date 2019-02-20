package apps

import (
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/labstack/echo"
)

type (
	HTTPHandler interface {
		GetApps(c echo.Context) error
		GetAppByID(c echo.Context) error
		UpdateApp(c echo.Context) error
	}

	// httpHandler instance
	httpHandler struct {
		Service Service
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, s Service) HTTPHandler {
	return &httpHandler{
		Service: s,
	}
}

// GetApps returns all apps as JSON from the AppService
func (a *httpHandler) GetApps(c echo.Context) error {
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

// GetAppByID returns an app with the ID in JSON format from the AppService
// TODO: Implement
//GetAppById returns a app with the ID in JSON format from the AppService
func (a *httpHandler) GetAppByID(c echo.Context) error {
	apps, err := a.Service.GetApps()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apps)
}

// UpdateApp returns an app updated with the ID in JSON format from the AppService
// TODO: Implement
//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *httpHandler) UpdateApp(c echo.Context) error {
	apps, err := a.Service.GetApps()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apps)
}
