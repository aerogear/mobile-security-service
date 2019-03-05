package apps

import (
	"database/sql"
	"fmt"
	"github.com/aerogear/mobile-security-service/pkg/models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
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
	COALESCE(COUNT(DISTINCT d.id),0) as num_of_current_installs
	FROM app as a LEFT JOIN version as v on a.app_id = v.app_id 
	LEFT JOIN device as d on v.id = d.version_id 
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
		if err = rows.Scan(&a.ID, &a.AppID, &a.AppName, &a.NumOfDeployedVersions, &a.NumOfAppLaunches, &a.NumOfCurrentInstalls); err != nil {
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
	rows, err := a.db.Query(`
	SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches,
	COALESCE(COUNT(DISTINCT d.id),0) as num_of_current_installs
	FROM version as v LEFT JOIN device as d on v.id = d.version_id
	WHERE v.app_id = $1 
	GROUP BY v.id;`, id)

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
		var disabledMessage sql.NullString
		if err = rows.Scan(&v.ID, &v.Version, &v.AppID, &v.Disabled, &disabledMessage, &v.NumOfCurrentInstalls, &v.NumOfAppLaunches); err != nil {
			log.Error(err)
		}

		v.DisabledMessage = disabledMessage.String
		versions = append(versions, v)
	}

	if len(versions) == 0 {
		return nil, models.ErrNotFound
	}

	return &versions, nil
}

// GetAppByID retrieves an app by id from the database
func (a *appsPostgreSQLRepository) GetAppByID(ID string) (*models.App, error) {
	var app models.App

	sqlStatement := `SELECT id,app_id,app_name FROM app WHERE id=$1;`
	row := a.db.QueryRow(sqlStatement, ID)
	err := row.Scan(&app.ID, &app.AppID, &app.AppName)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &app, nil

}

// UpdateAppVersions all versions sent
func (a *appsPostgreSQLRepository) UpdateAppVersions(versions []models.Version) error {

	for i := 0; i < len(versions); i++ {

		// Update Version
		_, err := a.db.Exec(`
		UPDATE version
		SET disabled_message=$1,disabled=$2
		WHERE ID=$3;`, versions[i].DisabledMessage, versions[i].Disabled, versions[i].ID)

		if err != nil {
			fmt.Print(err)
			log.Error(err)
			return err
		}
	}

	return nil
}

func (a *appsPostgreSQLRepository) DisableAllAppVersionsByAppID(appID string, message string) error {

	// Update Version
	_, err := a.db.Exec(`
		UPDATE version
		SET disabled_message=$1,disabled=True
		WHERE app_id=$2;`, message, appID)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *appsPostgreSQLRepository) DeleteAppByAppID(appId string) error {

	_, err := a.db.Exec(`
		UPDATE app
		SET deleted_at=$1
		WHERE app_id=$2;`, time.Now(), appId)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *appsPostgreSQLRepository) CreateApp(id, appId, name string) error {

	// Update Version
	_, err := a.db.Exec(`INSERT INTO app (id, app_id, app_name) VALUES ($1,$2,$3)`, id, appId, name)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *appsPostgreSQLRepository) GetAppByAppID(appID string) (*models.App, error) {

	var app models.App

	sqlStatement := `SELECT id,app_id,app_name FROM app WHERE app_id=$1;`
	row := a.db.QueryRow(sqlStatement, appID)
	err := row.Scan(&app.ID, &app.AppID, &app.AppName)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &app, nil
}

func (a *appsPostgreSQLRepository) UnDeleteAppByAppID(appId string) error {

	_, err := a.db.Exec(`
		UPDATE app
		SET deleted_at=NULL
		WHERE app_id=$1;`, appId)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
