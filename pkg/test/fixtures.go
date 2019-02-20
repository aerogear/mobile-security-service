package test

import (
	"database/sql"

	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/sirupsen/logrus"
)

// Seed the database with some sample data
func seedDatabase(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO app
			(id, app_id, app_name, deleted_at)
		VALUES 
			('1b9e7a5f-af7c-4055-b488-72f2b5f72266', 'com.aerogear.foobar', 'Foobar', NULL),
			('ae2da1f5-a9c4-4305-84bc-80da683fbc36', 'com.test.app1', 'App One', '2019-02-18 14:36:35'),
			('0890506c-3dd1-43ad-8a09-21a4111a65a6', 'com.aerogear.testapp', 'Test App', NULL);

		INSERT INTO version
			(id, version, app_id, disabled, disabled_message, num_of_app_launches, num_of_clients)
		VALUES 
			('f6fe70a3-8c99-429c-8c77-a2efa7d0b458', '1', 'com.aerogear.testapp', FALSE, '', 5000, 100),
    		('9bc87235-6bcb-40ab-993c-8722d86e2201', '1.1', 'com.aerogear.testapp', TRUE, 'Please contact an administrator', 1000, 59),
    		('def3c38b-5765-4041-a8e1-b2b60d58bece', '1', 'com.test.app1', FALSE, '', 10000, 200);`)

	if err != nil {
		logrus.Println(err)
	}
}

// GetMockAppList returns some dummy apps
func GetMockAppList() []models.App {
	apps := []models.App{
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc025",
			AppID:   "com.aerogear.mobile_app_one",
			AppName: "Mobile App One",
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
	return apps
}

// GetMockAppVersionList returns some dummy app versions
func GetMockAppVersionList() []models.Version {
	versions := []models.Version{
		models.Version{
			ID:               "55ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_one",
			DisabledMessage:  "Please contact an administrator",
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

	return versions
}
