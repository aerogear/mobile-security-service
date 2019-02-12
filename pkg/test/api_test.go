// +build integration

package test

import (
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"
	"github.com/aerogear/mobile-security-service/pkg/web/router"
)

func setUpTestServer() *httptest.Server {
	config := config.Get()
	dbConn, _ := db.Connect(config.DB.ConnectionString, config.DB.MaxConnections)
	e := router.NewRouter(config)

	apiRoutePrefix := config.ApiRoutePrefix
	apiGroup := e.Group(apiRoutePrefix)

	// App handler setup
	appsPostgreSQLRepository := apps.NewPostgreSQLRepository(dbConn)
	appsService := apps.NewService(appsPostgreSQLRepository)
	appsHandler := apps.NewHTTPHandler(e, appsService)

	// Setup routes
	router.SetAppRoutes(apiGroup, appsHandler)

	return httptest.NewServer(e)
}

func TestHTTPHandler_GetAppsEndpoint(t *testing.T) {
	server := setUpTestServer()

	tests := []struct {
		name       string
		wantStatus int
	}{
		{
			name:       "GetApps() should return a 200 status code with an array of data",
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := http.Get(server.URL + "/api/apps")

			if err != nil {
				t.Errorf("Got an unexpected error during GET request to /apps")
			}
			if res.StatusCode != tt.wantStatus {
				t.Errorf("HTTPHandler.GetApps() statusCode = %v, wantStatus %v", res.StatusCode, tt.wantStatus)
			}
		})
	}
}
