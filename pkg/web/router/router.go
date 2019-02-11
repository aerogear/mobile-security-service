package router

import (
	"github.com/aerogear/mobile-security-service/pkg/config"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/aerogear/mobile-security-service/pkg/web/middleware"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (v *RequestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewRouter(config config.Config) *echo.Echo {
	router := echo.New()

	middleware.Init(router, config)

	router.Validator = &RequestValidator{validator: validator.New()}
	return router
}

func SetAppRoutes(r *echo.Group, appsHandler *apps.HTTPHandler) {
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
	// operationId: getAppById
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
	r.GET("/apps/{id}", appsHandler.GetAppById) // TODO: Implement correctly the call of the method passing the parameters

	// swagger:operation PUT /apps/{id} App
	//
	// Update a single app using the app id, including updating version information
	// ---
	// summary: Update app by id
	// operationId: updateApp
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: The id for the app that needs to be updated.
	//   required: true
	//   type: string
	// - name: body
	//   in: body
	//   description: Updated app object
	//   required: true
	//   schema:
	//     $ref: '#/definitions/App'
	// responses:
	//   200:
	//     description: successful operation
	//     schema:
	//       $ref: '#/definitions/App'
	//   400:
	//     description: Invalid app supplied
	//   404:
	//     description: App not found
	r.PUT("/apps/{id}", appsHandler.UpdateApp) // TODO: Implement correctly the call of the method passing the parameters

	// swagger:operation POST /init appInitResponse
	//
	// Capture metrics from device and return if the app version they are using is disabled and has a set disabled message
	// ---
	// summary: Init call from sdk
	// operationId: initAppFromDevice
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   description: Updated app object
	//   required: true
	//   schema:
	//     $ref: '#/definitions/AppInit'
	// responses:
	//   200:
	//     description: successful operation
	//     schema:
	//       $ref: '#/definitions/AppInitResponse'
	//   400:
	//     description: Invalid appId supplied
	//   404:
	//     description: App not found
	r.POST("/init", appsHandler.GetApps) // TODO: Implement the func which will be used here and its call

}
