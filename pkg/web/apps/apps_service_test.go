package apps

import (
	"github.com/labstack/echo"
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"

	"github.com/aerogear/mobile-security-service/pkg/models"
)

func Test_appsService_GetApps(t *testing.T) {
	c := config.Get()
	dbConn, err := db.Connect(c.DB.ConnectionString, c.DB.MaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %v", err.Error())
	}

	postgresRepository := NewPostgreSQLRepository(dbConn)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		want    *[]models.App
		wantErr bool
	}{
		{
			name: "appsService.GetApps() should return a list of apps",
			want: &[]models.App{
				models.App{
					ID:                    "a0874c82-2b7f-11e9-b210-d663bd873d93",
					AppID:                 "com.aerogear.app1",
					AppName:               "app1",
					NumOfDeployedVersions: 1,
					NumOfAppLaunches:      1,
					NumOfClients:          1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(postgresRepository)

			got, err := a.GetApps()
			if (err != nil) != tt.wantErr {
				t.Errorf("appsService.GetApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsService.GetApps() = %v, want %v", got, tt.want)
			}
		})
	}
}
