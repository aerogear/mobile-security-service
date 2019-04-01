package apps

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	getAppsQueryString = `SELECT a.id,a.app_id,a.app_name,
	COALESCE\(COUNT\(DISTINCT v.id\),0\) as num_of_deployed_versions,
	COALESCE\(SUM\(DISTINCT v.num_of_app_launches\),0\) as num_of_app_launches,
	COALESCE\(COUNT\(DISTINCT d.id\),0\) as num_of_current_installs
	FROM app as a LEFT JOIN version as v on a.app_id = v.app_id 
	LEFT JOIN device as d on v.id = d.version_id 
	WHERE a.deleted_at IS NULL 
	GROUP BY a.id;`

	getAppVersionsQueryString = `SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches, v.last_launched_at,
	COALESCE\(COUNT\(DISTINCT d.id\),0\) as num_of_current_installs
	FROM version as v LEFT JOIN device as d on v.id = d.version_id
	WHERE v.app_id = \$1 
	GROUP BY v.id;`

	GetActiveAppByIDQueryString = `SELECT id,app_id,app_name FROM app WHERE deleted_at IS NULL AND id=\$1;`

	GetActiveAppByAppIDQueryString = `SELECT id,app_id,app_name FROM app WHERE app_id=\$1;`

	getUpdateAppVersionsQueryString = `UPDATE version
		SET disabled_message=\$1,disabled=\$2
		WHERE ID=\$3`

	getDeleteAppByIDQueryString = `UPDATE app
		SET deleted_at=\$1
		WHERE id=\$2;`

	getDisableAllAppVersionsByAppIDQueryString = `UPDATE version
		SET disabled_message=\$1,disabled=True
		WHERE app_id=\$2;`

	getUnDeleteAppByAppIDQueryString = `UPDATE app
		SET deleted_at=NULL
		WHERE app_id=\$1;`

	getUpdateAppNameByAppIDQueryString = `UPDATE app
		SET app_name=\$1
		WHERE app_id=\$2;`

	getDeviceByDeviceIDQuery = `SELECT id,version_id,app_id,device_id,device_type,device_version
	FROM devicei
	WHERE device_id = \$1;`

	getDeviceByDeviceIDAndAppIDQuery = `SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
	FROM device as d
	WHERE d.device_id = \$1 AND d.app_id = \$2;`

	getVersionByAppIDAndVersion = `SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches, v.last_launched_at
	FROM version as v
	WHERE v.app_id = \$1 AND v.version = \$2;`

	getDeviceByVersionAndAppIDQuery = `SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
		FROM device as d
		WHERE d.app_id = \$1 AND d.device_version = \$2;`

	GetActiveAppByAppIDQuery = `SELECT id,app_id,app_name FROM app WHERE app_id=\$1 AND deleted_at IS NULL;`

	GetAppByAppIDQuery = `SELECT id,app_id,app_name,deleted_at FROM app WHERE app_id=\$1;`

	upsertVersionWithAppLaunchesAndLastLaunchedStatement = `INSERT INTO version as v \(id, version, app_id, disabled, disabled_message, last_launched_at\)
		VALUES\(\$1, \$2, \$3, \$4, \$5, NOW\(\)\)
		ON CONFLICT \(id\)
		DO UPDATE
		SET num_of_app_launches = v\.num_of_app_launches \+ 1,
		last_launched_at = NOW\(\);`

	insertDeviceOrUpdateVersionIDStatement = `INSERT INTO device\(id,version_id,app_id,device_id,device_type,device_version\)
		VALUES\(\$1, \$2, \$3, \$4, \$5, \$6\)
		ON CONFLICT \(id\)
		DO UPDATE
		SET version_id = \$2, device_version = \$6`
)

func Test_appsPostgreSQLRepository_GetApps_WillReturnTwoApps(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := helpers.GetMockAppList()

	cols := []string{"id", "app_id", "app_name", "deleted_at"}

	timestamp := "2019-02-15T09:38:33+00:00"

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName, timestamp)

	// Insert 2 apps which are not soft deleted
	rows := sqlmock.NewRows(cols).AddRow(mockApps[1].ID, mockApps[1].AppID, mockApps[1].AppName, "").AddRow(mockApps[2].ID, mockApps[2].AppID, mockApps[2].AppName, "")

	// We should expected to get back only the apps which are not soft deleted
	mock.ExpectQuery(getAppsQueryString).WillReturnRows(rows)
	a := NewPostgreSQLRepository(db)

	apps, err := a.GetApps()

	if err != nil {
		t.Fatalf("Got error trying to get apps from database: %v", err)
	}

	if len(*apps) != 2 {
		t.Fatalf("Expected 2 apps to be returned from the database, got %v", len(*apps))
	}
}

func Test_appsPostgreSQLRepository_GetApps_WillReturnNoApps(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	timestamp := "2019-02-15T09:38:33+00:00"

	mockApps := helpers.GetMockAppList()

	// Insert 3 apps which are soft deleted
	sqlmock.NewRows([]string{"id", "app_id", "app_name", "deleted_at"}).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName, timestamp).AddRow(mockApps[1].ID, mockApps[1].AppID, mockApps[1].AppName, timestamp).AddRow(mockApps[2].ID, mockApps[2].AppID, mockApps[2].AppName, timestamp)

	// We should expected 0 apps
	mock.ExpectQuery(getAppsQueryString).WillReturnRows(&sqlmock.Rows{})
	a := NewPostgreSQLRepository(db)

	apps, err := a.GetApps()

	if err != nil && err != models.ErrNotFound {
		t.Fatalf("Expected ErrNotFound error to be returned from database, got %v", err)
	}

	if apps != nil && len(*apps) != 0 {
		t.Fatalf("Expected 0 apps to be returned from the database, got %v", len(*apps))
	}
}

func Test_appsPostgreSQLRepository_GetAppVersionsByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_one"

	mockVersions := helpers.GetMockAppVersionList()

	// Add the expected return data to the mock database
	rows := sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, mockVersions[0].Disabled, mockVersions[0].DisabledMessage, mockVersions[0].NumOfAppLaunches).
		AddRow(mockVersions[1].ID, mockVersions[1].Version, mockVersions[1].AppID, mockVersions[1].Disabled, mockVersions[1].DisabledMessage, mockVersions[1].NumOfAppLaunches)

	// Add some extra rows
	sqlmock.NewRows(cols).
		AddRow(mockVersions[2].ID, mockVersions[0].Version, mockVersions[2].AppID, mockVersions[2].Disabled, mockVersions[2].DisabledMessage, mockVersions[2].NumOfAppLaunches)

	mock.ExpectQuery(``).WithArgs(appID).WillReturnRows(rows)

	a := NewPostgreSQLRepository(db)

	versions, err := a.GetAppVersionsByAppID(appID)

	if err != nil {
		t.Fatalf("Got error trying to get apps from database: %v", err)
	}

	if len(*versions) != 2 {
		t.Fatalf("Expected 2 apps to be returned from the database, got %v", len(*versions))
	}
}

func Test_appsPostgreSQLRepository_GetAppVersions_WillReturnNoAppVersions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockVersions := helpers.GetMockAppVersionList()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_ten"

	// Add the expected return data to the mock database
	sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, mockVersions[0].Disabled, mockVersions[0].DisabledMessage, mockVersions[0].NumOfAppLaunches).
		AddRow(mockVersions[1].ID, mockVersions[1].Version, mockVersions[1].AppID, mockVersions[1].Disabled, mockVersions[1].DisabledMessage, mockVersions[1].NumOfAppLaunches)

	mock.ExpectQuery(getAppVersionsQueryString).WithArgs(appID).WillReturnRows(&sqlmock.Rows{})

	a := NewPostgreSQLRepository(db)

	versions, err := a.GetAppVersionsByAppID(appID)

	if err != nil && err != models.ErrNotFound {
		t.Fatalf("Expected ErrNotFound error to be returned from database, got %v", err)
	}

	if versions != nil && len(*versions) != 0 {
		t.Fatalf("Expected 0 apps to be returned from the database, got %v", len(*versions))
	}
}

func Test_appsPostgreSQLRepository_GetActiveAppByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()
	mockApps := helpers.GetMockAppList()
	cols := []string{"id", "app_id", "app_name", "deleted_at"}
	cols2 := []string{"id", "app_id", "app_name"}

	timestamp := "2019-02-15T09:38:33+00:00"

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName, timestamp)

	// Insert 2 apps which are not soft deleted
	rows := sqlmock.NewRows(cols2).AddRow(mockApps[1].ID, mockApps[1].AppID, mockApps[1].AppName).AddRow(mockApps[2].ID, mockApps[2].AppID, mockApps[2].AppName)

	tests := []struct {
		name      string
		ID        string
		row       *sqlmock.Rows
		cols      []string
		timestamp string
		wantErr   bool
	}{
		{
			name: "Get app by id should return an app",
			ID:   "7f89ce49-a736-459e-9110-e52d049fc026",
		},
		{
			name:    "Get app by id using an valid id format should return an error",
			ID:      "3489ce49-a736-459e-9110-e52d049fc030",
			wantErr: true,
		},
		{
			name:    "Get app by id using id that is soft deleted",
			ID:      "7f89ce49-a736-459e-9110-e52d049fc025",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		var app *models.App
		// We should expected to get back only the app with matching id
		mock.ExpectQuery(GetActiveAppByIDQueryString).WithArgs(tt.ID).WillReturnRows(rows)
		a := NewPostgreSQLRepository(db)

		app, err = a.GetActiveAppByID(tt.ID)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to get app from database: %v", err)
		}

		if app == &mockApps[0] {
			t.Fatalf("Expected an app to be returned from the database,want %v, got %v", &mockApps[0], app)
		}
	}
}

func Test_appsPostgreSQLRepository_DisableAllAppVersionsByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockVersions := helpers.GetMockAppVersionList()

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_one"
	msg := "disable"

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	for i := 0; i < len(mockVersions); i++ {
		sqlmock.NewRows(cols).
			AddRow(mockVersions[i].ID, mockVersions[i].Version, mockVersions[i].AppID, true, mockVersions[i].DisabledMessage, mockVersions[i].NumOfAppLaunches)

	}

	mock.ExpectExec(getDisableAllAppVersionsByAppIDQueryString).WithArgs(msg, appID).WillReturnResult(sqlmock.NewResult(0, 3))

	a := NewPostgreSQLRepository(db)

	if err = a.DisableAllAppVersionsByAppID(appID, msg); err != nil {
		t.Errorf("error was not expected while updating all versions: %s", err)
	}
}

// Error is expected when an invalid APP_ID is send
func Test_appsPostgreSQLRepository_DisableAllAppVersionsByAppID_ReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockVersions := helpers.GetMockAppVersionList()

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_one"
	msg := "disable"

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	for i := 0; i < len(mockVersions); i++ {
		sqlmock.NewRows(cols).
			AddRow(mockVersions[i].ID, mockVersions[i].Version, mockVersions[i].AppID, true, mockVersions[i].DisabledMessage, mockVersions[i].NumOfAppLaunches)

	}

	mock.ExpectExec(getDisableAllAppVersionsByAppIDQueryString).WithArgs(msg, "").WillReturnResult(sqlmock.NewResult(0, 3))

	a := NewPostgreSQLRepository(db)

	if err = a.DisableAllAppVersionsByAppID(appID, msg); err == nil {
		t.Errorf("error was expected while updating all versions: %s", err)
	}
}

func Test_appsPostgreSQLRepository_UpdateVersions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockVersions := helpers.GetMockAppVersionList()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, true, mockVersions[0].DisabledMessage, mockVersions[0].NumOfAppLaunches)

	// App ID which we expect to return the versions for
	id := "55ebd387-9c68-4137-a367-a12025cc2cdb"
	msg := "Please contact an administrator"
	status := true

	mock.ExpectExec(getUpdateAppVersionsQueryString).WithArgs(msg, status, id).WillReturnResult(sqlmock.NewResult(0, 3))

	a := NewPostgreSQLRepository(db)

	input := []models.Version{mockVersions[0]}
	input[0].Disabled = true
	if err = a.UpdateAppVersions(input); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func Test_appsPostgreSQLRepository_UpdateVersions_ReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockVersions := helpers.GetMockAppVersionList()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches"}

	sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, true, mockVersions[0].DisabledMessage, mockVersions[0].NumOfAppLaunches)

	// App ID which we expect to return the versions for
	id := "invalid"
	msg := "Please contact an administrator"
	status := true

	mock.ExpectExec(getUpdateAppVersionsQueryString).WithArgs(msg, status, id).WillReturnResult(sqlmock.NewResult(0, 3))

	a := NewPostgreSQLRepository(db)

	input := []models.Version{mockVersions[0]}
	input[0].Disabled = true
	if err = a.UpdateAppVersions(input); err == nil {
		t.Errorf("error was expected while updating stats: %s", err)
	}
}

type AnyTimestamp struct{}

func (e AnyTimestamp) Match(v driver.Value) bool {
	return true
}

func Test_appsPostgreSQLRepository_DeleteAppByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := helpers.GetMockAppList()

	cols := []string{"id", "app_id", "app_name"}
	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName)

	// We should expected to get back only the apps which are not soft deleted
	mock.ExpectExec(getDeleteAppByIDQueryString).WithArgs(AnyTimestamp{}, mockApps[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
	a := NewPostgreSQLRepository(db)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name: "Should update the deleted_at of an app with a value with success",
			id:   mockApps[0].ID,
		},
		{
			name:    "Should return an error when try to update the deleted_at of an app with a value",
			id:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {

		err = a.DeleteAppById(tt.id)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to update the deleted_at of an app with an value: %v", err)
		}

		if err == nil && tt.wantErr {
			t.Fatalf("Not get expected error when was trying to update the deleted_at of an app with an value")
		}
	}
}

func Test_appsPostgreSQLRepository_UnDeleteAppByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := helpers.GetMockApp()

	cols := []string{"id", "app_id", "app_name"}

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps.ID, mockApps.AppID, mockApps.AppName)

	// We should expected to get back only the apps which are not soft deleted
	mock.ExpectExec(getUnDeleteAppByAppIDQueryString).WithArgs(mockApps.AppID).WillReturnResult(sqlmock.NewResult(0, 1))
	a := NewPostgreSQLRepository(db)

	tests := []struct {
		name    string
		appId   string
		wantErr bool
	}{
		{
			name:  "Should return success when set deleted_at as NULL for an valid appID",
			appId: helpers.GetMockApp().AppID,
		},
		{
			name:    "Should return error when set deleted_at as NULL for an valid appID",
			appId:   helpers.GetMockApp().ID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		err = a.UnDeleteAppByAppID(tt.appId)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to get app from database: %v", err)
		}

		if err == nil && tt.wantErr {
			t.Fatalf("Not get expected error when was trying to create an app")
		}
	}
}

func Test_appsPostgreSQLRepository_UpdateAppNameByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := helpers.GetMockApp()

	cols := []string{"id", "app_id", "app_name"}

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps.ID, mockApps.AppID, mockApps.AppName)

	// We should expected to get back only the apps which are not soft deleted
	mock.ExpectExec(getUpdateAppNameByAppIDQueryString).WithArgs(mockApps.AppName, mockApps.AppID).WillReturnResult(sqlmock.NewResult(0, 1))
	a := NewPostgreSQLRepository(db)

	tests := []struct {
		name    string
		appId   string
		appName string
		wantErr bool
	}{
		{
			name:    "Should return success when set app_name for an valid appID",
			appId:   helpers.GetMockApp().AppID,
			appName: helpers.GetMockApp().AppName,
		},
		{
			name:    "Should return error when set app_name for an invalid appID",
			appId:   "",
			appName: helpers.GetMockApp().AppName,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		err = a.UpdateAppNameByAppID(tt.appId, tt.appName)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to get app from database: %v", err)
		}

		if err == nil && tt.wantErr {
			t.Fatalf("Not get expected error when was trying to create an app")
		}
	}
}

func Test_appsPostgreSQLRepository_GetDeviceByDeviceIDAndAppID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "version_id", "app_id", "device_id", "device_type", "device_version"}

	mockDevices := helpers.GetMockDevices(5)

	for _, d := range mockDevices {
		sqlmock.NewRows(cols).AddRow(d.ID, d.VersionID, d.AppID, d.DeviceID, d.DeviceType, d.DeviceVersion)
	}

	wantDevice := helpers.GetMockDevice()

	wantRow := sqlmock.NewRows(cols).AddRow(wantDevice.ID, wantDevice.VersionID, wantDevice.AppID, wantDevice.DeviceID, wantDevice.DeviceType, wantDevice.DeviceVersion)

	type args struct {
		deviceID string
		appID    string
	}
	type want struct {
		device *models.Device
		err    error
		rows   *sqlmock.Rows
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should return a device",
			args: args{
				deviceID: wantDevice.DeviceID,
				appID:    wantDevice.AppID,
			},
			want: want{
				device: wantDevice,
				err:    nil,
				rows:   wantRow,
			},
		},
		{
			name: "Should return error when invalid device ID supplied",
			args: args{
				deviceID: uuid.New().String(),
				appID:    wantDevice.AppID,
			},
			want: want{
				device: nil,
				err:    models.ErrNotFound,
				rows:   &sqlmock.Rows{},
			},
		},
		{
			name: "Should return an error when invalid app ID supplied",
			args: args{
				deviceID: wantDevice.DeviceID,
				appID:    uuid.New().String(),
			},
			want: want{
				device: nil,
				err:    models.ErrNotFound,
				rows:   &sqlmock.Rows{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPostgreSQLRepository(db)

			mock.ExpectQuery(getDeviceByDeviceIDAndAppIDQuery).WithArgs(tt.args.deviceID, tt.args.appID).WillReturnRows(tt.want.rows)

			got, err := repo.GetDeviceByDeviceIDAndAppID(tt.args.deviceID, tt.args.appID)

			if !reflect.DeepEqual(got, tt.want.device) {
				t.Errorf("appsPostgreSQLRepository.GetDeviceByDeviceIDAndAppID() = %v, want %v", got, tt.want.device)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetDeviceByDeviceIDAndAppID() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_GetDeviceByVersionAndAppID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "version_id", "app_id", "device_id", "device_type", "device_version"}

	mockDevices := helpers.GetMockDevices(5)

	for _, d := range mockDevices {
		sqlmock.NewRows(cols).AddRow(d.ID, d.VersionID, d.AppID, d.DeviceID, d.DeviceType, d.DeviceVersion)
	}

	wantDevice := helpers.GetMockDevice()

	wantRow := sqlmock.NewRows(cols).AddRow(wantDevice.ID, wantDevice.VersionID, wantDevice.AppID, wantDevice.DeviceID, wantDevice.DeviceType, wantDevice.DeviceVersion)

	type args struct {
		version string
		appID   string
	}
	type want struct {
		device *models.Device
		err    error
		rows   *sqlmock.Rows
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should return a device",
			args: args{
				version: wantDevice.Version,
				appID:   wantDevice.AppID,
			},
			want: want{
				device: wantDevice,
				err:    nil,
				rows:   wantRow,
			},
		},
		{
			name: "Should return error when invalid version is supplied",
			args: args{
				version: "0.0",
				appID:   wantDevice.AppID,
			},
			want: want{
				device: nil,
				err:    models.ErrNotFound,
				rows:   &sqlmock.Rows{},
			},
		},
		{
			name: "Should return an error when invalid app ID supplied",
			args: args{
				version: wantDevice.Version,
				appID:   uuid.New().String(),
			},
			want: want{
				device: nil,
				err:    models.ErrNotFound,
				rows:   &sqlmock.Rows{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPostgreSQLRepository(db)

			mock.ExpectQuery(getDeviceByVersionAndAppIDQuery).WithArgs(tt.args.version, tt.args.appID).WillReturnRows(tt.want.rows)

			got, err := repo.GetDeviceByVersionAndAppID(tt.args.appID, tt.args.version)

			if !reflect.DeepEqual(got, tt.want.device) {
				t.Errorf("appsPostgreSQLRepository.GetDeviceByVersionAndAppID() = %v, want %v", got, tt.want.device)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetDeviceByVersionAndAppID() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_GetActiveAppByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "app_id", "app_name"}

	mockApps := helpers.GetMockAppList()

	for _, a := range mockApps {
		sqlmock.NewRows(cols).AddRow(a.ID, a.AppID, a.AppName)
	}

	wantApp := helpers.GetMockApp()

	wantRow := sqlmock.NewRows(cols).AddRow(wantApp.ID, wantApp.AppID, wantApp.AppName)

	type args struct {
		appID string
	}
	type want struct {
		app  *models.App
		err  error
		rows *sqlmock.Rows
	}

	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Should return an app",
			args: args{
				appID: wantApp.AppID,
			},
			want: want{
				app:  wantApp,
				err:  nil,
				rows: wantRow,
			},
		},
		{
			name: "Should not return an app given invalid app ID",
			args: args{
				appID: uuid.New().String(),
			},
			want: want{
				app:  nil,
				err:  models.ErrNotFound,
				rows: &sqlmock.Rows{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery(GetActiveAppByAppIDQuery).WithArgs(tt.args.appID).WillReturnRows(tt.want.rows)

			repo := NewPostgreSQLRepository(db)
			got, err := repo.GetActiveAppByAppID(tt.args.appID)

			if (got != nil) && !reflect.DeepEqual(got.AppID, tt.want.app.AppID) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() = %v, want %v", got.AppID, tt.want.app.ID)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_GetActiveAppByAppID_ShouldReturnNoAppWhenAppIsSoftDeleted(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "app_id", "app_name", "deleted_at"}

	app := helpers.GetMockApp()

	sqlmock.NewRows(cols).AddRow(app.ID, app.AppID, app.AppName, "2019-02-15T09:38:33+00:00")

	type args struct {
		appID string
	}
	type want struct {
		app  *models.App
		err  error
		rows *sqlmock.Rows
	}

	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Should return ErrNotFound",
			args: args{
				appID: app.AppID,
			},
			want: want{
				app:  nil,
				err:  models.ErrNotFound,
				rows: &sqlmock.Rows{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery(GetActiveAppByAppIDQuery).WithArgs(tt.args.appID).WillReturnRows(tt.want.rows)

			repo := NewPostgreSQLRepository(db)
			got, err := repo.GetActiveAppByAppID(tt.args.appID)

			if (got != nil) && !reflect.DeepEqual(got.AppID, tt.want.app.AppID) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() = %v, want %v", got.AppID, tt.want.app.ID)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_GetAppByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "app_id", "app_name", "deleted_at"}

	mockApps := helpers.GetMockAppList()

	for _, a := range mockApps {
		sqlmock.NewRows(cols).AddRow(a.ID, a.AppID, a.AppName, a.DeletedAt)
	}

	wantApp := helpers.GetMockApp()

	wantRow := sqlmock.NewRows(cols).AddRow(wantApp.ID, wantApp.AppID, wantApp.AppName, wantApp.DeletedAt)

	type args struct {
		appID string
	}
	type want struct {
		app  *models.App
		err  error
		rows *sqlmock.Rows
	}

	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Should return an app",
			args: args{
				appID: wantApp.AppID,
			},
			want: want{
				app:  wantApp,
				err:  nil,
				rows: wantRow,
			},
		},
		{
			name: "Should not return an app given invalid app ID",
			args: args{
				appID: uuid.New().String(),
			},
			want: want{
				app:  nil,
				err:  models.ErrNotFound,
				rows: &sqlmock.Rows{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery(GetAppByAppIDQuery).WithArgs(tt.args.appID).WillReturnRows(tt.want.rows)

			repo := NewPostgreSQLRepository(db)
			got, err := repo.GetAppByAppID(tt.args.appID)

			if (got != nil) && !reflect.DeepEqual(got.AppID, tt.want.app.AppID) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() = %v, want %v", got.AppID, tt.want.app.ID)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetActiveAppByAppID() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_UpsertVersionWithAppLaunchesAndLastLaunched(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches", "last_launched_at"}

	mockVersions := helpers.GetMockAppVersionList()

	for _, v := range mockVersions {
		sqlmock.NewRows(cols).
			AddRow(v.ID, v.Version, v.AppID, v.Disabled, v.DisabledMessage, v.NumOfAppLaunches, v.LastLaunchedAt)
	}

	type args struct {
		version *models.Version
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should update successfully",
			args: args{
				version: &mockVersions[0],
			},
			wantErr: nil,
		},
		{
			name: "Should insert successfully",
			args: args{
				version: helpers.GetMockVersion(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			version := tt.args.version

			mock.ExpectExec(upsertVersionWithAppLaunchesAndLastLaunchedStatement).WithArgs(version.ID, version.Version, version.AppID, version.Disabled, version.DisabledMessage).WillReturnResult(sqlmock.NewResult(0, 1))

			repo := NewPostgreSQLRepository(db)

			if err := repo.UpsertVersionWithAppLaunchesAndLastLaunched(version); err != tt.wantErr {
				t.Errorf("appsPostgreSQLRepository.UpsertVersionWithAppLaunchesAndLastLaunched() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_InsertDeviceOrUpdateVersionID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	cols := []string{"id", "version_id", "app_id", "device_id", "device_type", "device_version"}

	mockDevices := helpers.GetMockDevices(3)

	for _, d := range mockDevices {
		sqlmock.NewRows(cols).
			AddRow(d.ID, d.VersionID, d.AppID, d.DeviceID, d.DeviceType, d.DeviceVersion)
	}

	type args struct {
		device *models.Device
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should update successfully",
			args: args{
				device: &mockDevices[0],
			},
			wantErr: nil,
		},
		{
			name: "Should insert successfully",
			args: args{
				device: helpers.GetMockDevice(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			device := tt.args.device

			mock.ExpectExec(insertDeviceOrUpdateVersionIDStatement).WithArgs(device.ID, device.VersionID, device.AppID, device.DeviceID, device.DeviceType, device.DeviceVersion).WillReturnResult(sqlmock.NewResult(0, 1))

			repo := NewPostgreSQLRepository(db)

			if err := repo.InsertDeviceOrUpdateVersionID(*device); err != tt.wantErr {
				t.Errorf("appsPostgreSQLRepository.InsertDeviceOrUpdateVersionID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_GetVersionByAppIDAndVersion(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_app_launches", "last_launched_at"}

	mockVersionList := helpers.GetMockAppVersionList()

	for _, v := range mockVersionList {
		sqlmock.NewRows(cols).AddRow(v.ID, v.Version, v.AppID, v.Disabled, v.DisabledMessage, v.NumOfAppLaunches, v.LastLaunchedAt)
	}

	wantVersion := helpers.GetMockVersion()

	row := sqlmock.NewRows(cols).AddRow(wantVersion.ID, wantVersion.Version, wantVersion.AppID, wantVersion.Disabled, wantVersion.DisabledMessage, wantVersion.NumOfAppLaunches, wantVersion.LastLaunchedAt)

	type args struct {
		appID         string
		versionNumber string
	}
	type want struct {
		version *models.Version
		err     error
	}
	tests := []struct {
		name     string
		args     args
		wantRows *sqlmock.Rows
		want     want
	}{
		{
			name: "should return a version when valid app ID and version number supplied",
			args: args{
				appID:         wantVersion.AppID,
				versionNumber: wantVersion.Version,
			},
			wantRows: row,
			want: want{
				version: wantVersion,
				err:     nil,
			},
		},
		{
			name: "should return ErrNotFound when invalid paramters provided",
			args: args{
				appID:         uuid.New().String(),
				versionNumber: "100",
			},
			wantRows: &sqlmock.Rows{},
			want: want{
				version: nil,
				err:     models.ErrNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery(getVersionByAppIDAndVersion).WithArgs(tt.args.appID, tt.args.versionNumber).WillReturnRows(tt.wantRows)

			repo := NewPostgreSQLRepository(db)

			got, err := repo.GetVersionByAppIDAndVersion(tt.args.appID, tt.args.versionNumber)

			if !reflect.DeepEqual(got, tt.want.version) {
				t.Errorf("appsPostgreSQLRepository.GetVersionByAppIDAndVersion() = %v, want %v", got, tt.want.version)
			}

			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("appsPostgreSQLRepository.GetVersionByAppIDAndVersion() error = %v, wantErr %v", got, tt.want.err)
			}
		})
	}
}

func Test_appsPostgreSQLRepository_CreateApp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := helpers.GetMockApp()

	cols := []string{"id", "app_id", "app_name"}

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps.ID, mockApps.AppID, mockApps.AppName)

	// We should expected to get back only the apps which are not soft deleted
	mock.ExpectExec("INSERT INTO app").WithArgs(mockApps.ID, mockApps.AppID, mockApps.AppName).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgreSQLRepository(db)

	tests := []struct {
		name    string
		id      string
		appId   string
		nameApp string
		wantErr bool
	}{
		{
			name:    "Should create an app",
			id:      mockApps.ID,
			appId:   mockApps.AppID,
			nameApp: mockApps.AppName,
		},
		{
			name:    "Should return an error when try to create an app since the ID is invalid",
			id:      mockApps.AppID,
			appId:   mockApps.ID,
			nameApp: mockApps.AppName,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		err = a.CreateApp(tt.id, mockApps.AppID, mockApps.AppName)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to create a new app from database: %v", err)
		}

		if err == nil && tt.wantErr {
			t.Fatalf("Not get expected error when was trying to create an app")
		}
	}
}
