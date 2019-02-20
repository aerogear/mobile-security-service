// +build integration

package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/web/apps"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"
	"github.com/aerogear/mobile-security-service/pkg/web/router"
)

// Sets up a test server so we can connect to the endpoints through HTTP calls
func setupTestServer() *httptest.Server {
	config := config.Get()

	dbConn, _ := db.Connect(config.DB.ConnectionString, config.DB.MaxConnections)
	// set up tables
	db.Setup(dbConn)

	// seed the database with some sample data
	seedDatabase(dbConn)

	e := router.NewRouter(config)

	APIRoutePrefix := config.APIRoutePrefix
	apiGroup := e.Group(APIRoutePrefix)

	// App handler setup
	appsPostgreSQLRepository := apps.NewPostgreSQLRepository(dbConn)
	appsService := apps.NewService(appsPostgreSQLRepository)
	appsHandler := apps.NewHTTPHandler(e, appsService)

	// Setup routes
	router.SetAppRoutes(apiGroup, appsHandler)

	return httptest.NewServer(e)
}

func TestHTTPHandler_GetAppsEndpoint(t *testing.T) {
	server := setupTestServer()

	tests := []struct {
		name       string
		wantStatus int
	}{
		{
			name:       "GetApps() should return a 200 status code with an array of data",
			wantStatus: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := http.Get(server.URL + "/api/apps")

			if err != nil {
				t.Errorf("Got an unexpected error during GET request to /apps")
			}
			if res.StatusCode != tt.wantStatus {
				t.Errorf("httpHandler.GetApps() statusCode = %v, wantStatus %v", res.StatusCode, tt.wantStatus)
			}
		})
	}
}
