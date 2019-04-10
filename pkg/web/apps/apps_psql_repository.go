package apps

import (
	"database/sql"

	"time"

	"github.com/aerogear/mobile-security-service/pkg/models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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

// GetAppVersionsByAppID returns app app versions with the provided app ID
func (a *appsPostgreSQLRepository) GetAppVersionsByAppID(id string) (*[]models.Version, error) {
	rows, err := a.db.Query(`
	SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches, v.last_launched_at,
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
		if err = rows.Scan(&v.ID, &v.Version, &v.AppID, &v.Disabled, &disabledMessage, &v.NumOfCurrentInstalls, &v.LastLaunchedAt, &v.NumOfAppLaunches); err != nil {
			log.Error(err)
		}

		v.DisabledMessage = disabledMessage.String
		versions = append(versions, v)
	}

	if len(versions) == 0 {
		return &versions, models.ErrNotFound
	}

	return &versions, nil
}

// GetActiveAppByID retrieves an app by id from the database
func (a *appsPostgreSQLRepository) GetActiveAppByID(ID string) (*models.App, error) {
	var app models.App

	sqlStatement := `SELECT id,app_id,app_name FROM app WHERE deleted_at IS NULL AND id=$1;`
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

// GetVersionByAppIDAndVersion gets a version by its app ID and version number
func (a *appsPostgreSQLRepository) GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error) {
	version := models.Version{}

	sqlStatement := `
	SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches, v.last_launched_at
	FROM version as v
	WHERE v.app_id = $1 AND v.version = $2;`

	err := a.db.QueryRow(sqlStatement, appID, versionNumber).Scan(&version.ID, &version.Version, &version.AppID, &version.Disabled, &version.DisabledMessage, &version.NumOfAppLaunches, &version.LastLaunchedAt)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &version, nil
}

// GetDeviceByDeviceIDAndAppID returns a device by its device ID and app ID
func (a *appsPostgreSQLRepository) GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error) {
	device := models.Device{}

	sqlStatement := `
		SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
		FROM device as d
		WHERE d.device_id = $1 AND d.app_id = $2;`

	err := a.db.QueryRow(sqlStatement, deviceID, appID).Scan(&device.ID, &device.VersionID, &device.AppID, &device.DeviceID, &device.DeviceType, &device.DeviceVersion)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &device, nil
}

// GetDeviceByVersionAndAppID returns a device by its version number and app ID
func (a *appsPostgreSQLRepository) GetDeviceByVersionAndAppID(version string, appID string) (*models.Device, error) {
	device := models.Device{}

	sqlStatement := `
		SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
		FROM device as d
		WHERE d.app_id = $1 AND d.device_version = $2;`

	err := a.db.QueryRow(sqlStatement, appID, version).Scan(&device.ID, &device.VersionID, &device.AppID, &device.DeviceID, &device.DeviceType, &device.DeviceVersion)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &device, nil
}

// GetAppByID retrieves an app by id from the database
func (a *appsPostgreSQLRepository) GetAppByAppID(appID string) (*models.App, error) {
	app := models.App{}

	sqlStatement := `SELECT id,app_id,app_name,deleted_at FROM app WHERE app_id=$1;`

	var deletedAt sql.NullString
	err := a.db.QueryRow(sqlStatement, appID).Scan(&app.ID, &app.AppID, &app.AppName, &deletedAt)
	app.DeletedAt = deletedAt.String

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &app, nil
}

// GetActiveAppByID retrieves an app by id from the database where it is not soft deleted
func (a *appsPostgreSQLRepository) GetActiveAppByAppID(appID string) (*models.App, error) {
	app := models.App{}

	sqlStatement := `SELECT id,app_id,app_name FROM app WHERE app_id=$1 AND deleted_at IS NULL;`

	err := a.db.QueryRow(sqlStatement, appID).Scan(&app.ID, &app.AppID, &app.AppName)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &app, nil
}

// UpsertVersionWithAppLaunchesAndLastLaunched creates a new version row
// or increments the num_of_app_launches counter if the version already exists
func (a *appsPostgreSQLRepository) UpsertVersionWithAppLaunchesAndLastLaunched(version *models.Version) error {
	sqlStatement := `
		INSERT INTO version as v (id, version, app_id, disabled, disabled_message, last_launched_at)
		VALUES($1, $2, $3, $4, $5, NOW())
		ON CONFLICT (id)
		DO UPDATE
		SET num_of_app_launches = v.num_of_app_launches + 1,
		last_launched_at = NOW();`

	_, err := a.db.Exec(sqlStatement, version.ID, version.Version, version.AppID, version.Disabled, version.DisabledMessage)

	if err != nil {
		log.Error(err)
		return models.ErrInternalServerError
	}

	return nil
}

// InsertDeviceOrUpdateVersionID creates a new device row in the device table
func (a *appsPostgreSQLRepository) InsertDeviceOrUpdateVersionID(device models.Device) error {
	sqlStatement := `
		INSERT INTO device(id,version_id,app_id,device_id,device_type,device_version)
		VALUES($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id)
		DO UPDATE
		SET version_id = $2, device_version = $6`

	_, err := a.db.Exec(sqlStatement, device.ID, device.VersionID, device.AppID, device.DeviceID, device.DeviceType, device.DeviceVersion)

	if err != nil {
		log.Error(err)
		return models.ErrInternalServerError
	}

	return nil
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
			log.Error(err)
			return err
		}
	}

	return nil
}

// DisableAllAppVersionsAndSetDisabledMessageByAppID disables all app versions by its app ID
func (a *appsPostgreSQLRepository) DisableAllAppVersionsAndSetDisabledMessageByAppID(appID string, message string) error {

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

// DisableAllAppVersionsByAppID disables all app versions by its app ID
func (a *appsPostgreSQLRepository) DisableAllAppVersionsByAppID(appID string) error {

	// Update Version
	_, err := a.db.Exec(`
		UPDATE version
		SET disabled=True
		WHERE app_id=$1;`, appID)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *appsPostgreSQLRepository) DeleteAppById(id string) error {

	_, err := a.db.Exec(`
		UPDATE app
		SET deleted_at=$1
		WHERE id=$2;`, time.Now(), id)

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

func (a *appsPostgreSQLRepository) UpdateAppNameByID(id, name string) error {

	_, err := a.db.Exec(`
		UPDATE app
		SET app_name=$1
		WHERE id=$2;`, name, id)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
