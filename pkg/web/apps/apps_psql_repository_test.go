package apps

import (
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/helpers"

	"github.com/aerogear/mobile-security-service/pkg/models"
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

	getAppVersionsQueryString = `SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches,
	COALESCE\(COUNT\(DISTINCT d.id\),0\) as num_of_current_installs
	FROM version as v LEFT JOIN device as d on v.id = d.version_id
	WHERE v.app_id = \$1 
	GROUP BY v.id;`
	getAppByIDQueryString = `SELECT id,app_id,app_name FROM app WHERE id=\$1;`
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

func Test_appsPostgreSQLRepository_GetAppByID(t *testing.T) {
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
		ID      string
		wantErr bool
	}{
		{
			name: "Get app by id should return an app",
			ID:   "7f89ce49-a736-459e-9110-e52d049fc025",
		},
		{
			name:    "Get app by id using an valid id format should return an error",
			ID:      "3489ce49-a736-459e-9110-e52d049fc025",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		var app *models.App
		// We should expected to get back only the app with matching id
		mock.ExpectQuery(getAppByIDQueryString).WithArgs(tt.ID).WillReturnRows(row)
		a := NewPostgreSQLRepository(db)

		app, err = a.GetAppByID(tt.ID)

		if err != nil && !tt.wantErr {
			t.Fatalf("Got error trying to get app from database: %v", err)
		}

		if app == &mockApp[0] {
			t.Fatalf("Expected an app to be returned from the database,want %v, got %v", &mockApp[0], app)
		}
	}
}
