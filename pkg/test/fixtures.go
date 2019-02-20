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
			(id, version, app_id, disabled, disabled_message, num_of_app_launches)
		VALUES 
			('f6fe70a3-8c99-429c-8c77-a2efa7d0b458', '1', 'com.aerogear.testapp', FALSE, '', 5000),
    	('9bc87235-6bcb-40ab-993c-8722d86e2201', '1.1', 'com.aerogear.testapp', TRUE, 'Please contact an administrator', 1000),
			('def3c38b-5765-4041-a8e1-b2b60d58bece', '1', 'com.test.app1', FALSE, '', 10000);
				
		INSERT INTO device
			(id, version_id, app_id, device_id, device_type, device_version)
		VALUES 
			('d19feeb4-fb21-44e8-9990-473bf97a0a3f', 'f6fe70a3-8c99-429c-8c77-a2efa7d0b458', 'com.aerogear.testapp', 'a742f8b7-5e2f-43f3-a3c8-073da858420f', 'iOS', '10.2'),
			('00cb8957-db04-4ab6-8fd8-14b9fc516dbd', '9bc87235-6bcb-40ab-993c-8722d86e2201', 'com.aerogear.testapp', 'd1895cc1-28d7-4283-932d-8bcab9e4a461', 'Android', '3.2'),
			('e3b43b01-167b-48ef-8ff4-caf2e6613dee', '9bc87235-6bcb-40ab-993c-8722d86e2201', 'com.aerogear.testapp', 'feee7f81-0e33-4548-abbb-17a681c12f3b', 'Android', '4.1'),
			('ab411c3e-29f8-4e70-9ddc-8bafbba3fc4c', 'def3c38b-5765-4041-a8e1-b2b60d58bece', 'com.test.app1', '94da9833-093e-4f4c-9a93-b11600ce46b7', 'iOS', '2.0'),
			('a42a128a-dfb6-435c-8653-8f66ab3a5a1c', 'def3c38b-5765-4041-a8e1-b2b60d58bece', 'com.test.app1', '94132b0c-d7b1-4419-bcce-fc6760c59e3a', 'Android', '4.1');
	`)

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
			ID:                   "55ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:              "1.0",
			AppID:                "com.aerogear.mobile_app_one",
			DisabledMessage:      "Please contact an administrator",
			Disabled:             false,
			NumOfCurrentInstalls: 1,
			NumOfAppLaunches:     2,
		},
		models.Version{
			ID:                   "59ebd387-9c68-4137-a367-a12025cc1cdb",
			Version:              "1.1",
			AppID:                "com.aerogear.mobile_app_one",
			Disabled:             false,
			NumOfCurrentInstalls: 0,
			NumOfAppLaunches:     0,
		},
		models.Version{
			ID:                   "59dbd387-9c68-4137-a367-a12025cc2cdb",
			Version:              "1.0",
			AppID:                "com.aerogear.mobile_app_two",
			Disabled:             false,
			NumOfCurrentInstalls: 0,
			NumOfAppLaunches:     0,
		},
	}

	return versions
}
