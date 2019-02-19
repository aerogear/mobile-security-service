package test

import (
	"database/sql"
	"strings"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"
	"github.com/aerogear/mobile-security-service/models"
)

func SetUpDatabase() *sql.DB {
	c := config.Get()
	dbConn, err := db.Connect(c.DB.ConnectionString, c.DB.MaxConnections)

	if err != nil {
		panic("failed to connect to SQL database: " + err.Error())
	}

	if err := db.Setup(dbConn); err != nil {
		panic("failed to perform database setup: " + err.Error())
	}

	return dbConn
}

func trimBody(body string) string {
	return strings.TrimSpace(body)
}

func GetValidAppModelList() []models.App {
	return []models.App{
		models.App{
			ID:                    "a0874c82-2b7f-11e9-b210-d663bd873d93",
			AppID:                 "com.aerogear.app112392730",
			AppName:               "aerogear-mobile-app",
			NumOfDeployedVersions: 0,
			NumOfAppLaunches:      0,
			NumOfClients:          0,
		},
		models.App{
			ID:                    "a0874c82-2b7f-11e9-b210-d663bd873d93",
			AppID:                 "com.aerogear.app1234567890",
			AppName:               "aerogear-mobile-app-two",
			NumOfDeployedVersions: 30,
			NumOfAppLaunches:      10000,
			NumOfClients:          100,
		},
	}
}

func GetEmptyApp() models.App {
	return models.App{}
}
