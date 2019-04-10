package user

import (
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

const (
	USER_NAME_HEADER  = "X-Forwarded-User"
	USER_EMAIL_HEADER = "X-Forwarded-Email"
)

type (
	HTTPHandler interface {
		GetUser(c echo.Context) error
	}

	// httpHandler instance
	httpHandler struct{}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo) HTTPHandler {
	return &httpHandler{}
}

// GetUser returns the user if Oauth added headers are present
func (a *httpHandler) GetUser(c echo.Context) error {

	var user models.User

	//these headers values will be set by the openshift oauth-proxy
	if userNameHeader := c.Request().Header[USER_NAME_HEADER]; userNameHeader != nil && userNameHeader[0] != "" {
		user.Username = userNameHeader[0]
	} else {
		return httperrors.NotFound(c, "No User Found")
	}
	if userEmailHeader := c.Request().Header[USER_EMAIL_HEADER]; userEmailHeader != nil && userEmailHeader[0] != "" {
		user.Email = userEmailHeader[0]
	}

	return c.JSON(http.StatusOK, user)

}
