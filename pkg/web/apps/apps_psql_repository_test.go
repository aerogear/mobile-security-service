// +build integration

package apps

import (
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/db"

	"github.com/aerogear/mobile-security-service/models"
)

func Test_appsPostgreSQLRepository_GetApps(t *testing.T) {
	c := config.Get()
	dbConn, err := db.Connect(c.DB.ConnectionString, c.DB.MaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %v", err.Error())
	}

	tests := []struct {
		name    string
		want    *[]models.App
		wantErr bool
	}{
		{
			name: "appsPostgreSQLRepository.GetApps() should return a list of apps",
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
			a := NewPostgreSQLRepository(dbConn)
			got, err := a.GetApps()
			if (err != nil) != tt.wantErr {
				t.Errorf("appsPostgreSQLRepository.GetApps() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsPostgreSQLRepository.GetApps() = %v, want %v", got, tt.want)
			}
		})
	}
}
