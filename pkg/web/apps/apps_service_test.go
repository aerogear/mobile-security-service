package apps

import (
	"reflect"
	"testing"

	"github.com/google/uuid"

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
		GetActiveAppByAppIDFunc: func(appID string) (*models.App, error) {
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
		GetActiveAppByAppIDFunc: func(appID string) (*models.App, error) {
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
		GetActiveAppByAppIDFunc: func(appID string) (*models.App, error) {
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
		GetActiveAppByAppIDFunc: func(appID string) (*models.App, error) {
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

func Test_appsService_InitClientApp(t *testing.T) {
	validDevice := helpers.GetMockDevice()

	deviceWithoutDeviceIDAndAppID := *validDevice
	deviceWithoutDeviceIDAndAppID.DeviceID = ""
	deviceWithoutDeviceIDAndAppID.AppID = ""

	deviceWithNewVersion := *validDevice
	deviceWithNewVersion.Version = "10"

	deviceWithNewDeviceID := *validDevice
	deviceWithNewDeviceID.DeviceID = uuid.New().String()

	deviceWithNewDeviceVersion := *validDevice
	deviceWithNewDeviceVersion.DeviceVersion = "10"

	mockVersion := helpers.GetMockVersion()
	mockVersion.ID = validDevice.VersionID

	type fields struct {
		app     *models.App
		version *models.Version
		device  *models.Device
	}
	type args struct {
		deviceInfo *models.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Version
		wantErr error
	}{
		{
			name:    "InitClient() should return ErrNotFound when app not found",
			wantErr: models.ErrNotFound,
			want:    nil,
			args: args{
				deviceInfo: validDevice,
			},
			fields: fields{
				app: nil,
			},
		},
		{
			name:    "InitClient() should return init client data",
			wantErr: nil,
			args: args{
				deviceInfo: validDevice,
			},
			fields: fields{
				app: &models.App{
					ID:      uuid.New().String(),
					AppID:   "com.aerogear.testapp",
					AppName: "Test App",
				},
				version: mockVersion,
				device:  validDevice,
			},
		},
		{
			name:    "InitClient() should create a new device when device with device ID and app ID not found",
			wantErr: nil,
			args: args{
				deviceInfo: &deviceWithoutDeviceIDAndAppID,
			},
			fields: fields{
				app: &models.App{
					ID:      uuid.New().String(),
					AppID:   "com.aerogear.testapp",
					AppName: "Test App",
				},
				version: mockVersion,
				device:  validDevice,
			},
		},
		{
			name:    "Insert a new version and return init data when new version number is not found",
			wantErr: nil,
			args: args{
				deviceInfo: &deviceWithNewVersion,
			},
			fields: fields{
				app: &models.App{
					ID:      uuid.New().String(),
					AppID:   "com.aerogear.testapp",
					AppName: "Test App",
				},
				version: mockVersion,
				device:  &deviceWithNewVersion,
			},
		},
		{
			name:    "should update device when new device ID found",
			wantErr: nil,
			args: args{
				deviceInfo: &deviceWithNewDeviceID,
			},
			fields: fields{
				app: &models.App{
					ID:      uuid.New().String(),
					AppID:   "com.aerogear.testapp",
					AppName: "Test App",
				},
				version: mockVersion,
				device:  validDevice,
			},
		},
		{
			name:    "should update device when new device version found",
			wantErr: nil,
			args: args{
				deviceInfo: &deviceWithNewDeviceVersion,
			},
			fields: fields{
				app: &models.App{
					ID:      uuid.New().String(),
					AppID:   "com.aerogear.testapp",
					AppName: "Test App",
				},
				version: mockVersion,
				device:  validDevice,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockedRepository := &RepositoryMock{
				GetActiveAppByAppIDFunc: func(appID string) (*models.App, error) {
					if tt.fields.app == nil {
						return nil, models.ErrNotFound
					}

					return tt.fields.app, nil
				},
				GetVersionByAppIDAndVersionFunc: func(appID string, version string) (*models.Version, error) {
					if (tt.fields.version != nil) && (tt.fields.version.AppID != appID || tt.fields.version.Version != version) {
						return nil, models.ErrNotFound
					}

					return tt.fields.version, nil
				},
				UpsertVersionWithAppLaunchesAndLastLaunchedFunc: func(version *models.Version) error {
					if version == nil {
						return models.ErrDatabaseError
					}
					return nil
				},
				GetDeviceByVersionAndAppIDFunc: func(version string, appID string) (*models.Device, error) {
					if (tt.fields.version != nil) && (tt.fields.version.Version != version || tt.args.deviceInfo.AppID != appID) {
						return nil, models.ErrNotFound
					}

					return tt.args.deviceInfo, nil
				},
				GetDeviceByDeviceIDAndAppIDFunc: func(deviceID string, appID string) (*models.Device, error) {
					if tt.fields.device == nil {
						return nil, models.ErrNotFound
					}

					if tt.fields.device.DeviceID != deviceID || tt.fields.device.AppID != appID {
						return nil, models.ErrNotFound
					}

					return tt.fields.device, nil
				},
				InsertDeviceOrUpdateVersionIDFunc: func(device models.Device) error {
					if &device == nil {
						return models.ErrDatabaseError
					}

					return nil
				},
			}

			service := NewService(mockedRepository)

			got, err := service.InitClientApp(tt.args.deviceInfo)

			if err == nil {
				tt.want = &models.Version{
					ID:              tt.fields.version.ID,
					Version:         tt.fields.version.Version,
					AppID:           tt.fields.version.AppID,
					Disabled:        tt.fields.version.Disabled,
					DisabledMessage: tt.fields.version.DisabledMessage,
				}
			}

			if err != tt.wantErr {
				t.Errorf("appsService.InitClientApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If a new version was not created, expect to return an exact device match
			if tt.fields.version != nil && (tt.fields.version.Version == tt.args.deviceInfo.Version && tt.fields.version.AppID == tt.args.deviceInfo.AppID) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsService.InitClientApp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
