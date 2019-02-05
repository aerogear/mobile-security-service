package router

import (
	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/aerogear/mobile-security-service/pkg/web/middleware"

	"github.com/labstack/echo"

	"gopkg.in/go-playground/validator.v9"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (v *RequestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewRouter(config config.Config) *echo.Echo {
	router := echo.New()

	middleware.Init(router, config)

	router.Validator = &RequestValidator{validator: validator.New()}
	return router
}

func SetAppRoutes(r *echo.Group, appsHandler *apps.HTTPHandler) {
	r.GET("/apps", appsHandler.GetApps)
}