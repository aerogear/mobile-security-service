package middleware

import (
	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init - Initialise custom middleware
func Init(e *echo.Echo, c config.Config) {
	e.Use(corsWithConfig(c)) // CORS
}

// corsWithConfig defines custom CORS rules for this server
func corsWithConfig(c config.Config) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     c.CORS.AllowOrigins,
		AllowCredentials: c.CORS.AllowCredentials,
	})
}
