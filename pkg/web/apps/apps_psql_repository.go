package apps

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/aerogear/mobile-security-service/pkg/models"
)

type (
	// PostgreSQLRepository interface defines the methods to be implemented
	appsPostgreSQLRepository struct {
		db *sql.DB
	}
)

// NewPostgreSQLRepository creates a new instance of appsPostgreSQLRepository
func NewPostgreSQLRepository(db *sql.DB) Repository {
	return &appsPostgreSQLRepository{db}
}

// GetApps retrieves all apps from the database
func (a *appsPostgreSQLRepository) GetApps() (*[]models.App, error) {
	rows, err := a.db.Query(`
	SELECT a.id,a.app_id,a.app_name,
	COALESCE(COUNT(DISTINCT v.id),0) as num_of_deployed_versions,
	COALESCE(SUM(DISTINCT v.num_of_app_launches),0) as num_of_app_launches,
	COALESCE(SUM(DISTINCT v.num_of_clients),0) as num_of_clients
	FROM app as a LEFT JOIN version as v on a.app_id = v.app_id 
	WHERE a.deleted_at IS NULL 
	GROUP BY a.id;`)

	if err != nil {
		log.Error(err)
		return nil, models.ErrInternalServerError
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	apps := []models.App{}
	for rows.Next() {
		var a models.App
		if err = rows.Scan(&a.ID, &a.AppID, &a.AppName, &a.NumOfDeployedVersions, &a.NumOfAppLaunches, &a.NumOfClients); err != nil {
			log.Error(err)
		}

		apps = append(apps, a)
	}

	if len(apps) == 0 {
		return nil, models.ErrNotFound
	}

	return &apps, nil
}

func (a *appsPostgreSQLRepository) GetAppVersionsByAppID(id string) (*[]models.Version, error) {
	rows, err := a.db.Query(`SELECT id,version,app_id,disabled,disabled_message,num_of_clients,num_of_app_launches FROM version WHERE app_id = $1`, id)

	if err != nil {
		log.Error(err)
		return nil, models.ErrInternalServerError
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	versions := []models.Version{}

	// iterate over the rows and add the data to the array of versions
	for rows.Next() {
		var v models.Version
		if err = rows.Scan(&v.ID, &v.Version, &v.AppID, &v.Disabled, &v.DisabledMessage, &v.NumOfClients, &v.NumOfAppLaunches); err != nil {
			log.Error(err)
		}

		versions = append(versions, v)
	}

	if len(versions) == 0 {
		return nil, models.ErrNotFound
	}

	return &versions, nil
}
