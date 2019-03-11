package checks

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/labstack/echo"
	"net/http"
)

type (
	// HTTPHandler instance
	HTTPHandler struct {
		appsService apps.Service
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, a apps.Service) *HTTPHandler {
	return &HTTPHandler{
		appsService: a,
	}
}

//Check if the server is alive - Liveness Probe
func (a *HTTPHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

//Check if the server is able to receive requests - Readiness
func (a *HTTPHandler) Healthz(c echo.Context) error {
	// TODO: Create an specific service to check if it is readiness
	_, err := a.appsService.GetApps()
	if err != nil {
		if err == models.ErrNotFound {
			return c.JSON(http.StatusOK, "OK")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
