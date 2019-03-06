package apps

import (
	"encoding/json"
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

type (
	HTTPHandler interface {
		GetApps(c echo.Context) error
		GetActiveAppByID(c echo.Context) error
		UpdateAppVersions(c echo.Context) error
		DisableAllAppVersionsByAppID(c echo.Context) error
		InitClientApp(c echo.Context) error
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

// GetActiveAppByID returns apps by id as JSON from the AppService
func (a *httpHandler) GetActiveAppByID(c echo.Context) error {

	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	apps, err := a.Service.GetActiveAppByID(id)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}
	return c.JSON(http.StatusOK, apps)

}

//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *httpHandler) UpdateAppVersions(c echo.Context) error {
	// Validations
	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	versions := []models.Version{}
	errV := json.NewDecoder(c.Request().Body).Decode(&versions)

	// check if the data sent is in the correct format
	if errV != nil {
		return httperrors.BadRequest(c, "Invalid data")
	}

	// Check if versions were sent all the body is empty
	if len(versions) == 0 {
		return httperrors.BadRequest(c, "No version(s) was sent.")
	}

	// Call service
	errUpdate := a.Service.UpdateAppVersions(versions)
	if errUpdate != nil {
		return httperrors.GetHTTPResponseFromErr(c, errUpdate)
	}

	return c.JSON(http.StatusOK, "")
}

//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *httpHandler) DisableAllAppVersionsByAppID(c echo.Context) error {
	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	// Transform the body request in the version struct
	ver := models.Version{}
	errV := json.NewDecoder(c.Request().Body).Decode(&ver)

	// check if the data sent is in the correct format
	if errV != nil {
		return httperrors.BadRequest(c, "Invalid data")
	}

	err := a.Service.DisableAllAppVersionsByAppID(id, ver.DisabledMessage)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, "")

}

// InitClientApp stores device information and returns if the app version is disabled
func (a *httpHandler) InitClientApp(c echo.Context) error {
	// Transform the body request in the version struct
	device := models.Device{}
	err := json.NewDecoder(c.Request().Body).Decode(&device)

	// check if the data sent is in the correct format
	if err != nil {
		return httperrors.BadRequest(c, "Invalid data")
	}

	if err := a.validateInitCallDataProvided(device, c); err != nil {
		return err
	}

	response, err := a.Service.InitClientApp(&device)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, response)

}

func (a *httpHandler) validateInitCallDataProvided(device models.Device, c echo.Context) error {

	if !helpers.IsValidUUID(device.DeviceID) {
		return httperrors.BadRequest(c, "Invalid DeviceID supplied")
	}

	if !helpers.IsValidUUID(device.VersionID) {
		return httperrors.BadRequest(c, "Invalid VersionID supplied")
	}

	if !helpers.IsValidUUID(device.AppID) {
		return httperrors.BadRequest(c, "Invalid AppID supplied")
	}

	if device.Version == "" {
		return httperrors.BadRequest(c, "Invalid Version supplied")
	}
	return nil
}
