package router

import (
	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/aerogear/mobile-security-service/pkg/web/checks"
	"github.com/aerogear/mobile-security-service/pkg/web/initclient"
	"github.com/aerogear/mobile-security-service/pkg/web/middleware"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type requestValidator struct {
	validator *validator.Validate
}

// Validate validates structs and individual fields based on tags.
func (v *requestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// NewRouter creates and returns a new instance of the Echo framework
func NewRouter(config config.Config) *echo.Echo {
	router := echo.New()

	middleware.Init(router, config)

	router.Validator = &requestValidator{validator: validator.New()}
	return router
}

// SetAppRoutes binds the route address to their handler functions
func SetAppRoutes(r *echo.Group, appsHandler apps.HTTPHandler) {
	// swagger:operation GET /apps App
	//
	// Returns root level information for all apps
	// ---
	// summary: Retrieve list of apps
	// operationId: getApps
	// produces:
	// - application/json
	// parameters: []
	// responses:
	//   200:
	//     description: successful operation
	//     schema:
	//       $ref: '#/definitions/App'
	r.GET("/apps", appsHandler.GetApps)

	// swagger:operation GET /apps/{id} App
	//
	// Retrieve all information for a single app including all child information
	// ---
	// summary: Get app by id
	// operationId: GetActiveAppByID
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: The id for the app that needs to be fetched.
	//   required: true
	//   type: string
	// responses:
	//   200:
	//     description: successful operation
	//     schema:
	//       $ref: '#/definitions/App'
	//   400:
	//     description: Invalid id supplied
	//   404:
	//     description: App not found
	r.GET("/apps/:id", appsHandler.GetActiveAppByID)
	// swagger:operation PUT /apps/:id/versions Version
	//
	// Update all versions informed of an app using the app id, including updating version information
	// ---
	// summary: Update 1 or more versions of an app
	// operationId: UpdateAppVersions
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: The id for the app that will have its versions updated
	//   required: true
	//   type: string
	// - name: body
	//   in: body
	//   description: Updated 1 or more versions of an app
	//   required: true
	//   schema:
	//     $ref: '#/definitions/Version'
	// responses:
	//   200:
	//     description: successful update
	//   400:
	//     description: Invalid app and/or versions supplied
	//   404:
	//     description: App not found
	r.PUT("/apps/:id/versions", appsHandler.UpdateAppVersions)

	// swagger:operation POST /apps/:id/versions/disable Version
	//
	// Disable all versions of an app
	// ---
	// summary: Disable all versions of an app
	// operationId: updateApp
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: The id for the app that will have all its versions updated
	//   required: true
	//   type: string
	// - name: body
	//   in: body
	//   description:
	//   required: true
	//   schema:
	//     $ref: '#/definitions/Version'
	// responses:
	//   200:
	//     description: successful update
	//   400:
	//     description: Invalid app supplied
	//   404:
	//     description: App not found
	r.POST("/apps/:id/versions/disable", appsHandler.DisableAllAppVersionsByAppID)

	// Binding/Re-Binding an app
	// ---
	// summary:
	// - Create an new app
	// - Update deleted_at as NULL to re-binding the app if the app informed be disabled
	// - Update the name if the name informed is not "" and is different of the name stored in the database for the appId
	// operationId: BindingApp
	// produces:
	// - application/json
	// parameters:
	// - name: appId
	//   in: path
	//   description: The appId of the app
	//   required: true
	//   type: string
	// - name: body
	//   in: body
	//   description:
	//   required: true
	//   schema:
	//     $ref: '#/definitions/App'
	// responses:
	//   200:
	//     description: successful operation
	//   400:
	//     description: Invalid id supplied
	r.POST("/apps/bind/:appId", appsHandler.BindingApp)
}

func SetInitRoutes(r *echo.Group, initHandler *initclient.HTTPHandler) {
	// swagger:operation POST /init Device
	//
	// Capture metrics from device and return if the app version they are using is disabled and has a set disabled message
	// ---
	// summary: Init call from SDK
	// operationId: initAppFromDevice
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   description: Updated app object
	//   required: true
	//   schema:
	//     $ref: '#/definitions/Version'
	// responses:
	//   200:
	//     description: successful operation
	//   400:
	//     description: Invalid id supplied
	//   404:
	//     description: Data not found
	r.POST("/init", initHandler.InitClientApp)
}

func SetChecksRouter(r *echo.Group, handler *checks.HTTPHandler) {
	// swagger:operation GET /ping Status
	//
	// Check the status of the REST SERVICE API
	// ---
	// summary: Check if the server is running
	// operationId: status
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: successful operation
	r.GET("/ping", handler.Ping)

	// swagger:operation GET /healthz Status
	//
	// Check the health of the REST SERVICE API
	// ---
	// summary: Check if the server can receive requests
	// operationId: health
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: successful operation
	//   500:
	//     description: Internal Server Error
	r.GET("/healthz", handler.Healthz)
}
