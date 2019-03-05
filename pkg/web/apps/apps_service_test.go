package apps

import (
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
)

var (
	mockRepositoryWithSuccessResults = &RepositoryMock{
		GetActiveAppByIDFunc: func(ID string) (*models.App, error) {
			return helpers.GetMockApp(), nil
		},
		GetAppVersionsByAppIDFunc: func(ID string) (*[]models.Version, error) {
			res := helpers.GetMockAppVersionList()
			return &res, nil
		},
		GetAppsFunc: func() (*[]models.App, error) {
			apps := helpers.GetMockAppList()
			return &apps, nil
		},
		UpdateAppVersionsFunc: func(versions []models.Version) error {
			return nil
		},
		DisableAllAppVersionsByAppIDFunc: func(id string, message string) error {
			return nil
		},
		DeleteAppByAppIDFunc: func(appId string) error {
			return nil
		},
		CreateAppFunc: func(id string, appId string, name string) error {
			return nil
		},
		GetAppByAppIDFunc: func(appID string) (*models.App, error) {
			return nil, models.ErrNotFound
		},
		UnDeleteAppByAppIDFunc: func(appID string) error {
			return nil
		},
	}

	mockRepositoryError = &RepositoryMock{
		GetActiveAppByIDFunc: func(ID string) (*models.App, error) {
			return nil, models.ErrNotFound
		},
		GetAppVersionsByAppIDFunc: func(ID string) (*[]models.Version, error) {
			return nil, models.ErrNotFound
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return nil, models.ErrNotFound
		},
		UpdateAppVersionsFunc: func(versions []models.Version) error {
			return models.ErrNotFound
		},
		DisableAllAppVersionsByAppIDFunc: func(id string, message string) error {
			return models.ErrNotFound
		},
		DeleteAppByAppIDFunc: func(appId string) error {
			return models.ErrNotFound
		},
		CreateAppFunc: func(id string, appId string, name string) error {
			return models.ErrConflict
		},
		GetAppByAppIDFunc: func(appID string) (*models.App, error) {
			return helpers.GetMockApp(), nil
		},
		UnDeleteAppByAppIDFunc: func(appID string) error {
			return nil
		},
	}
)

func Test_appsService_GetApps(t *testing.T) {
	apps := helpers.GetMockAppList()

	type fields struct {
		repository Repository
	}
	tests := []struct {
		name     string
		fields   fields
		id       string
		want     *[]models.App
		wantErr  error
		mockRepo RepositoryMock
	}{
		{
			name:     "Get all apps should return success",
			want:     &apps,
			mockRepo: *mockRepositoryWithSuccessResults,
		},
		{
			name:     "Get all apps should return error when apps are not found",
			want:     &apps,
			wantErr:  models.ErrNotFound,
			mockRepo: *mockRepositoryError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(&tt.mockRepo)
			got, err := a.GetApps()
			if (err != nil) && tt.wantErr == nil {
				t.Errorf("appsService.GetApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil) && (!reflect.DeepEqual(got, tt.want)) {
				t.Errorf("appsService.GetApps() = %v, want %v", got, tt.want)
			}
			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.GetApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_appsService_GetActiveAppByID(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name     string
		fields   fields
		id       string
		want     *models.App
		wantErr  error
		mockRepo RepositoryMock
	}{
		{
			name:     "Get app by id",
			id:       "7f89ce49-a736-459e-9110-e52d049fc025",
			want:     helpers.GetMockApp(),
			mockRepo: *mockRepositoryWithSuccessResults,
		},
		{
			name:     "Return an error as no file app found",
			id:       "7f89ce49-a736-459e-9110-e52d049fcerr",
			want:     nil,
			wantErr:  models.ErrNotFound,
			mockRepo: *mockRepositoryError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(&tt.mockRepo)
			got, err := a.GetActiveAppByID(tt.id)
			if (err != nil) && tt.wantErr == nil {
				t.Errorf("appsService.GetActiveAppByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsService.GetActiveAppByID() = %v, want %v", got, tt.want)
			}
			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.GetActiveAppByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_appsService_DisableAllAppVersionsByAppID(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name    string
		fields  fields
		id      string
		msg     string
		wantErr error
		repo    RepositoryMock
	}{
		{
			name: "Disable all app versions",
			id:   "7f89ce49-a736-459e-9110-e52d049fc025",
			msg:  "disable",
			repo: *mockRepositoryWithSuccessResults,
		},
		{
			name:    "Return error to update the version",
			id:      "invalid",
			repo:    *mockRepositoryError,
			wantErr: models.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(mockRepositoryWithSuccessResults)
			err := a.DisableAllAppVersionsByAppID(tt.id, tt.msg)
			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.DisableAllAppVersionsByAppID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_appsService_UpdateAppVersions(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name     string
		fields   fields
		versions []models.Version
		wantErr  error
		repo     RepositoryMock
	}{
		{
			name:     "Update versions",
			versions: helpers.GetMockAppVersionList(),
			repo:     *mockRepositoryWithSuccessResults,
		},
		{
			name:    "Return error to update the version",
			repo:    *mockRepositoryError,
			wantErr: models.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(&tt.repo)
			err := a.UpdateAppVersions(tt.versions)

			if (err != nil) && tt.wantErr == nil {
				t.Errorf("appsService.DisableAllAppVersionsByAppID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.DisableAllAppVersionsByAppID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_appsService_UnbindingAppByAppID(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name    string
		fields  fields
		appId   string
		wantErr error
		repo    RepositoryMock
	}{
		{
			name:  "Should unbinding app by app_id",
			appId: helpers.GetMockApp().AppID,
			repo:  *mockRepositoryWithSuccessResults,
		},
		{
			name:    "Should return an error to unbinding app",
			appId:   helpers.GetMockApp().AppID,
			repo:    *mockRepositoryError,
			wantErr: models.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(&tt.repo)
			err := a.UnbindingAppByAppID(tt.appId)

			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.DeleteAppByAppID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_appsService_BindingApp(t *testing.T) {
	// make and configure a mocked Repository
	mockRepositoryWithNewBindingSuccessResults := &RepositoryMock{
		UnDeleteAppByAppIDFunc: func(appID string) error {
			return nil
		},
		GetAppByAppIDFunc: func(appID string) (*models.App, error) {
			return nil, models.ErrNotFound
		},
		CreateAppFunc: func(id string, appId string, name string) error {
			return nil
		},
		GetActiveAppByIDFunc: func(ID string) (*models.App, error) {
			return helpers.GetMockApp(), nil
		},
	}

	// make and configure a mocked Repository
	mockRepositoryWithErrorToGetAppId := &RepositoryMock{
		UnDeleteAppByAppIDFunc: func(appID string) error {
			return nil
		},
		GetAppByAppIDFunc: func(appID string) (*models.App, error) {
			return nil, models.ErrInternalServerError
		},
		CreateAppFunc: func(id string, appId string, name string) error {
			return nil
		},
	}

	type fields struct {
		repository Repository
	}
	tests := []struct {
		name    string
		fields  fields
		appId   string
		nameApp string
		wantErr error
		repo    RepositoryMock
	}{
		{
			name:    "Should binding an new app by app_id and name",
			appId:   helpers.GetMockApp().AppID,
			nameApp: helpers.GetMockApp().AppName,
			repo:    *mockRepositoryWithNewBindingSuccessResults,
		},
		{
			name:    "Should re-binding an app by app_id and name",
			appId:   helpers.GetMockApp().AppID,
			nameApp: helpers.GetMockApp().AppName,
			repo:    *mockRepositoryWithSuccessResults,
		},
		{
			name:    "Should return error to binding an new app by app_id and name",
			appId:   helpers.GetMockApp().AppID,
			nameApp: helpers.GetMockApp().AppName,
			repo:    *mockRepositoryError,
			wantErr: models.ErrConflict,
		},
		{
			name:    "Should return error to binding an new app by app_id",
			appId:   helpers.GetMockApp().AppID,
			nameApp: helpers.GetMockApp().AppName,
			repo:    *mockRepositoryWithErrorToGetAppId,
			wantErr: models.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewService(&tt.repo)
			err := a.BindingAppByApp(tt.appId, tt.name)

			if (err != nil) && (tt.wantErr != err || tt.wantErr == nil) {
				t.Errorf("appsService.BindingAppByApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
