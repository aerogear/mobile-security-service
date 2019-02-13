package apps

import (
	"database/sql"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/models"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	getAppsQueryString = `SELECT a.id,a.app_id,a.app_name,
	COALESCE\(COUNT\(DISTINCT v.id\),0\) as num_of_deployed_versions,
	COALESCE\(SUM\(DISTINCT v.num_of_app_launches\),0\) as num_of_app_launches,
	COALESCE\(SUM\(DISTINCT v.num_of_clients\),0\) as num_of_clients
	FROM app as a LEFT JOIN version as v on a.app_id = v.app_id 
	WHERE a.deleted_at IS NULL 
	GROUP BY a.id;`
	getAppVersionsQueryString = `SELECT id,version,app_id,disabled,disabled_message,num_of_clients,num_of_app_launches FROM version WHERE app_id = \$1`
)

func Test_appsPostgreSQLRepository_GetApps_WillReturnTwoApps(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error opening a stub database connection: %v", err)
	}

	defer db.Close()

	mockApps := []models.App{
		models.App{
			ID:        "7f89ce49-a736-459e-9110-e52d049fc025",
			AppID:     "com.aerogear.mobile_app_one",
			AppName:   "Mobile App One",
			DeletedAt: sql.NullString{String: "2011-01-01 00:00:00", Valid: true},
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc026",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Two",
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc027",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Three",
		},
	}

	cols := []string{"id", "app_id", "app_name", "deleted_at"}

	// Insert an app where the deleted_at column is set
	sqlmock.NewRows(cols).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName, mockApps[0].DeletedAt)

	// Insert 2 apps which are not soft deleted
	rows := sqlmock.NewRows(cols).AddRow(mockApps[1].ID, mockApps[1].AppID, mockApps[1].AppName, mockApps[1].DeletedAt).AddRow(mockApps[2].ID, mockApps[2].AppID, mockApps[2].AppName, mockApps[2].DeletedAt)

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

	mockApps := []models.App{
		models.App{
			ID:        "7f89ce49-a736-459e-9110-e52d049fc025",
			AppID:     "com.aerogear.mobile_app_one",
			AppName:   "Mobile App One",
			DeletedAt: sql.NullString{String: "2011-01-01 00:00:00", Valid: true},
		},
		models.App{
			ID:        "7f89ce49-a736-459e-9110-e52d049fc026",
			AppID:     "com.aerogear.mobile_app_three",
			AppName:   "Mobile App Two",
			DeletedAt: sql.NullString{String: "2011-01-01 00:00:00", Valid: true},
		},
		models.App{
			ID:        "7f89ce49-a736-459e-9110-e52d049fc027",
			AppID:     "com.aerogear.mobile_app_three",
			AppName:   "Mobile App Three",
			DeletedAt: sql.NullString{String: "2011-01-01 00:00:00", Valid: true},
		},
	}

	// Insert 3 apps which are soft deleted
	sqlmock.NewRows([]string{"id", "app_id", "app_name", "deleted_at"}).AddRow(mockApps[0].ID, mockApps[0].AppID, mockApps[0].AppName, mockApps[0].DeletedAt).AddRow(mockApps[1].ID, mockApps[1].AppID, mockApps[1].AppName, mockApps[1].DeletedAt).AddRow(mockApps[2].ID, mockApps[2].AppID, mockApps[2].AppName, mockApps[2].DeletedAt)

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

	mockVersions := []models.Version{
		models.Version{
			ID:               "55ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_one",
			Disabled:         false,
			NumOfClients:     0,
			NumOfAppLaunches: 0,
		},
		models.Version{
			ID:               "59ebd387-9c68-4137-a367-a12025cc1cdb",
			Version:          "1.1",
			AppID:            "com.aerogear.mobile_app_one",
			Disabled:         false,
			NumOfClients:     0,
			NumOfAppLaunches: 0,
		},
		models.Version{
			ID:               "59dbd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_two",
			Disabled:         false,
			NumOfClients:     0,
			NumOfAppLaunches: 0,
		},
	}

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_clients", "num_of_app_launches"}

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_one"

	// Add the expected return data to the mock database
	rows := sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, mockVersions[0].Disabled, mockVersions[0].DisabledMessage, mockVersions[0].NumOfClients, mockVersions[0].NumOfAppLaunches).
		AddRow(mockVersions[1].ID, mockVersions[1].Version, mockVersions[1].AppID, mockVersions[1].Disabled, mockVersions[1].DisabledMessage, mockVersions[1].NumOfClients, mockVersions[1].NumOfAppLaunches)

	// Add some extra rows
	sqlmock.NewRows(cols).
		AddRow(mockVersions[2].ID, mockVersions[0].Version, mockVersions[2].AppID, mockVersions[2].Disabled, mockVersions[2].DisabledMessage, mockVersions[2].NumOfClients, mockVersions[2].NumOfAppLaunches)

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

	mockVersions := []models.Version{
		models.Version{
			ID:               "59ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_one",
			Disabled:         false,
			NumOfClients:     0,
			NumOfAppLaunches: 0,
		},
		models.Version{
			ID:               "58ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.1",
			AppID:            "com.aerogear.mobile_app_one",
			Disabled:         false,
			NumOfClients:     0,
			NumOfAppLaunches: 0,
		},
	}

	cols := []string{"id", "version", "app_id", "disabled", "disabled_message", "num_of_clients", "num_of_app_launches"}

	// App ID which we expect to return the versions for
	appID := "com.aerogear.mobile_app_ten"

	// Add the expected return data to the mock database
	sqlmock.NewRows(cols).
		AddRow(mockVersions[0].ID, mockVersions[0].Version, mockVersions[0].AppID, mockVersions[0].Disabled, mockVersions[0].DisabledMessage, mockVersions[0].NumOfClients, mockVersions[0].NumOfAppLaunches).
		AddRow(mockVersions[1].ID, mockVersions[1].Version, mockVersions[1].AppID, mockVersions[1].Disabled, mockVersions[1].DisabledMessage, mockVersions[1].NumOfClients, mockVersions[1].NumOfAppLaunches)

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
