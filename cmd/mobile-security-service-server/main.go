package main

import (
	dotenv "github.com/joho/godotenv"

	"github.com/aerogear/mobile-security-service-server/pkg/config"
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
	e := echo.New()

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
