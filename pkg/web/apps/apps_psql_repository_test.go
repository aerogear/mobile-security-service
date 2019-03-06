package apps

import (
	"database/sql/driver"
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
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

	getAppVersionsQueryString = `SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches,
	COALESCE\(COUNT\(DISTINCT d.id\),0\) as num_of_current_installs
	FROM version as v LEFT JOIN device as d on v.id = d.version_id
	WHERE v.app_id = \$1 
	GROUP BY v.id;`

	GetActiveAppByIDQueryString = `SELECT id,app_id,app_name FROM app WHERE deleted_at IS NULL AND id=\$1;`

	getAppByAppIDQueryString = `SELECT id,app_id,app_name FROM app WHERE app_id=\$1;`

	getUpdateAppVersionsQueryString = `UPDATE version
		SET disabled_message=\$1,disabled=\$2
		WHERE ID=\$3`

	getDeleteAppByAppIDQueryString = `UPDATE app
		SET deleted_at=\$1
		WHERE app_id=\$2;`

	getDisableAllAppVersionsByAppIDQueryString = `UPDATE version
		SET disabled_message=\$1,disabled=True
		WHERE app_id=\$2;`

	getUnDeleteAppByAppIDQueryString = `UPDATE app
		SET deleted_at=NULL
		WHERE app_id=\$1;`
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

func Test_appsPostgreSQLRepository_DeleteAppByAppID(t *testing.T) {
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
	mock.ExpectExec(getDeleteAppByAppIDQueryString).WithArgs(AnyTimestamp{}, mockApps[0].AppID).WillReturnResult(sqlmock.NewResult(0, 1))
	a := NewPostgreSQLRepository(db)

	tests := []struct {
		name    string
		appId   string
		wantErr bool
	}{
		{
			name:  "Should update the deleted_at of an app with a value with success",
			appId: helpers.GetMockApp().AppID,
		},
		{
			name:    "Should return an error when try to update the deleted_at of an app with a value",
			appId:   helpers.GetMockApp().ID,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		err = a.DeleteAppByAppID(tt.appId)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to update the deleted_at of an app with an value: %v", err)
		}

		if err == nil && tt.wantErr {
			t.Fatalf("Not get expected error when was trying to update the deleted_at of an app with an value")
		}
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

func Test_appsPostgreSQLRepository_GetAppByAppID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()
	mockApp := helpers.GetMockAppList()
	cols := []string{"id", "app_id", "app_name"}
	// Insert app
	row := sqlmock.NewRows(cols).AddRow(mockApp[0].ID, mockApp[0].AppID, mockApp[0].AppName)

	tests := []struct {
		name    string
		appId   string
		wantErr bool
	}{
		{
			name:  "Get app by id should return an app",
			appId: helpers.GetMockApp().AppID,
		},
		{
			name:    "Get app by id using an valid id format should return an error",
			appId:   helpers.GetMockApp().ID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		var app *models.App
		// We should expected to get back only the app with matching id
		mock.ExpectQuery(getAppByAppIDQueryString).WithArgs(tt.appId).WillReturnRows(row)
		a := NewPostgreSQLRepository(db)

		app, err = a.GetAppByAppID(tt.appId)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to get app by id from database: %v", err)
		}

		if app == &mockApp[0] {
			t.Fatalf("Expected an app to be returned from the database,want %v, got %v", &mockApp[0], app)
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
