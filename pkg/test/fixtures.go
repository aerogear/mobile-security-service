package test

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Seed the database with some sample data
func seedDatabase(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO app
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
    		('def3c38b-5765-4041-a8e1-b2b60d58bece', '1', 'com.test.app1', FALSE, '', 10000, 200);
	`)

	if err != nil {
		logrus.Println(err)
	}
}
