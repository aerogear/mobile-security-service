package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"reflect"
	"testing"
)

func Test_appsService_GetApps(t *testing.T) {

	// mock data
	app := models.App{
		ID:                    "a0874c82-2b7f-11e9-b210-d663bd873d93",
		AppID:                 "com.aerogear.app1",
		AppName:               "app1",
		NumOfDeployedVersions: 1,
		NumOfAppLaunches:      1,
		NumOfClients:          1,
	}

	// make and configure a mocked Service
	mockedAppService := &AppServiceMock{
		GetAppsFunc: func() (*[]models.App, error) {
			return &[]models.App{
				app,
			}, nil
		},
	}

	tests := []struct {
		name    string
		want    *[]models.App
		wantErr bool
	}{
		{
			name: "appsService.GetApps() should return a list of apps",
			want: &[]models.App{
				app,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockedAppService.GetApps()
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
