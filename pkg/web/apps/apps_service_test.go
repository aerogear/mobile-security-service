package apps

import (
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_appsService_GetApps(t *testing.T) {
	apps := helpers.GetMockAppList()
	// make and configure a mocked Repository
	mockedRepository := &RepositoryMock{
		GetAppVersionsByAppIDFunc: func(id string) (*[]models.Version, error) {
			panic("mock out the GetAppVersionsByAppID method")
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return &apps, nil
		},
	}

	s := &appsService{mockedRepository}

	// Assertions
	got, err := s.GetApps()
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, &apps, got)
}

func Test_appsService_GetAppByID(t *testing.T) {
	// make and configure a mocked Repository
	app := helpers.GetMockApp()
	versions := helpers.GetMockAppVersionList()
	apps := helpers.GetMockAppList()
	mockedRepository := &RepositoryMock{
		GetAppByIDFunc: func(ID string) (*models.App, error) {
			if ID == app.ID {
				return app, nil
			}
			return nil, models.ErrNotFound
		},
		GetAppVersionsByAppIDFunc: func(id string) (*[]models.Version, error) {
			return &versions, nil
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return &apps, nil
		},
	}
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name    string
		fields  fields
		id      string
		want    *models.App
		wantErr bool
	}{
		{
			name:    "Get app by id",
			id:      "7f89ce49-a736-459e-9110-e52d049fc025",
			want:    app,
			wantErr: false,
		},
		{
			name:    "Return an error as no file app found",
			id:      "7f89ce49-a736-459e-9110-e52d049fcerr",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(mockedRepository)
			got, err := a.GetAppByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("appsService.GetAppByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsService.GetAppByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
