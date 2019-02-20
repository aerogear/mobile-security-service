package middleware

import (
	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Init - Initialise custom middleware
func Init(e *echo.Echo, c config.Config) {
	e.Use(corsWithConfig(c)) // CORS
	e.Use(Logger)
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  c.StaticFilesDir,
		HTML5: true,
		Index: "index.html",
		Skipper: func(context echo.Context) bool {
			// We don't want to return the SPA if any api/* is called, it should act like a normal API.
			return strings.HasPrefix(context.Request().URL.Path, c.APIRoutePrefix)
		},
		Browse: false,
	}))
}

// corsWithConfig defines custom CORS rules for this server
func corsWithConfig(c config.Config) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     c.CORS.AllowOrigins,
		AllowCredentials: c.CORS.AllowCredentials,
	})
}

// Logger logs information about all incoming requests
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		r := c.Request()

		log.WithFields(log.Fields{
			"method":        r.Method,
			"path":          r.RequestURI,
			"client_ip":     r.RemoteAddr,
			"response_time": time.Since(start),
		}).Info()

		return next(c)
	}
}
