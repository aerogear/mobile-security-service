package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	ApiRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "A counter for total HTTP requests",
		},
		[]string{"code", "method", "path"},
	)
	ApiRequestsFailureTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_failure_total",
			Help: "A counter for total HTTP request failures",
		},
		[]string{"code", "method", "path"},
	)
	ApiRequestsInFlight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "api_requests_in_flight",
		Help: "A gauge of concurrent requests",
	})
	ApiRequestsDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "api_requests_duration_seconds",
			Help: "HTTP request latencies in seconds.",
		},
		[]string{"method", "path"},
	)
)

// Init - Initialise custom middleware
func Init(e *echo.Echo, c config.Config) {
	prometheus.MustRegister(ApiRequestsTotal, ApiRequestsFailureTotal, ApiRequestsInFlight, ApiRequestsDuration)

	e.Use(corsWithConfig(c)) // CORS
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

// LogHTTPMetrics captures information about all incoming requests for Prometheus
func LogHTTPMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		ApiRequestsInFlight.Inc()

		// Execute the http handler by calling next(c)
		if err := next(c); err != nil {
			c.Error(err)
		}

		status := strconv.Itoa(res.Status)
		method := req.Method
		path := req.RequestURI
		elapsed := time.Since(start).Seconds()

		// Capture Prometheus metrics
		ApiRequestsTotal.WithLabelValues(status, method, path).Inc()
		ApiRequestsDuration.WithLabelValues(method, path).Observe(elapsed)
		ApiRequestsInFlight.Dec()

		// Log request details at Info level
		log.WithFields(log.Fields{
			"status":        status,
			"method":        req.Method,
			"path":          req.RequestURI,
			"client_ip":     req.RemoteAddr,
			"response_time": time.Since(start),
		}).Info()

		return nil
	}
}
