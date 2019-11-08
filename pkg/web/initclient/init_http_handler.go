package initclient

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/labstack/echo"
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

// InitClientApp stores device information and returns if the app version is disabled
func (h *HTTPHandler) InitClientApp(c echo.Context) error {
	deviceInfo := new(models.Device)

	if err := c.Bind(deviceInfo); err != nil {
		log.Info(err)
		return err
	}

	// Check the request body is valid
	if err := validateInitBody(deviceInfo); err != nil {
		log.Info(err)
		return httperrors.BadRequest(c, err.Error())
	}

	initResponse, err := h.appsService.InitClientApp(deviceInfo)

	// If no app has been found in the database, return a bad request to the client
	if err == models.ErrNotFound {
		return httperrors.BadRequest(c, "No bound app found for the sent App ID")
	}

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, initResponse)
}

// validateInitBody validates the properties of an init
// request and returns an error if any of them are missing
func validateInitBody(d *models.Device) error {
	if d.Version == "" {
		return errors.New("version property is required")
	}

	if d.AppID == "" {
		return errors.New("appId property is required")
	}

	if !helpers.IsValidUUID(d.DeviceID) {
		return errors.New("deviceId is invalid")
	}

	return nil
}
