// API for Mobile Security Service
//
// This is a sample mobile security service server.
//
//     Schemes: http, https
//     Version: 0.1.0
//     basePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"database/sql"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/aerogear/mobile-security-service/pkg/web/initclient"
	"github.com/aerogear/mobile-security-service/pkg/web/router"
	dotenv "github.com/joho/godotenv"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func init() {
	config := config.Get()

	initLogger(config.LogLevel, config.LogFormat)

	err := dotenv.Load()

	if err != nil {
		log.Info("No .env file found, using default values instead.")
	}
}

func main() {
	config := config.Get()

	e := router.NewRouter(config)

	db := connectDatabase(config)
	setupServer(e, config, db)

	// start webserver
	if err := e.Start(config.ListenAddress); err != nil {
		panic("failed to start" + err.Error())
	}
}

func initLogger(level, format string) {
	logLevel, err := log.ParseLevel(level)

	if err != nil {
		log.Fatalf("log level %v is not allowed. Must be one of [debug, info, warning, error, fatal, panic]", level)
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)

	switch format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{DisableColors: true})
	default:
		log.Fatalf("log format %v is not allowed. Must be one of [text, json]", format)
	}
}

// Make a connection to the PostgreSQL database
func connectDatabase(c config.Config) *sql.DB {
	dbConn, err := db.Connect(c.DB.ConnectionString, c.DB.MaxConnections)

	if err != nil {
		panic("failed to connect to SQL database: " + err.Error())
	}

	if err := db.Setup(dbConn); err != nil {
		panic("failed to perform database setup: " + err.Error())
	}

	return dbConn
}

// Invoke handlers, services and repositories here
func setupServer(e *echo.Echo, c config.Config, dbConn *sql.DB) {
	// Prefix api routes
	APIRoutePrefix := c.APIRoutePrefix
	apiGroup := e.Group(APIRoutePrefix)

	// App handler setup
	appsPostgreSQLRepository := apps.NewPostgreSQLRepository(dbConn)
	appsService := apps.NewService(appsPostgreSQLRepository)
	appsHandler := apps.NewHTTPHandler(e, appsService)

	// Setup app routes
	router.SetAppRoutes(apiGroup, appsHandler)

	// Initclient handler setup
	initclientHandler := initclient.NewHTTPHandler(e, appsService)

	// Setup initclient routes
	router.SetInitRoutes(apiGroup, initclientHandler)
}
