package httperrors

import (
	"encoding/json"

	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

var codes = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Unordered Collection",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	509: "Bandwidth Limit Exceeded",
	510: "Not Extended",
	511: "Network Authentication Required",
}

type errResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// BadRequest response code (400) indicates that the
// server could not understand the request due to invalid syntax.
func BadRequest(c echo.Context, message string) (e error) {
	return HTTPError(c, 400, message)
}

// Unauthorized response code (401) indicates that the
// request has not been applied because it lacks valid
// authentication credentials for the target resource.
func Unauthorized(c echo.Context, message string) (e error) {
	return HTTPError(c, 401, message)
}

// Forbidden response code (403) indicates that the
// server understood the request but refuses to authorize it.
func Forbidden(c echo.Context, message string) (e error) {
	return HTTPError(c, 403, message)
}

// NotFound response code (404) indicates that
// the server can't find the requested resource
func NotFound(c echo.Context, message string) (e error) {
	return HTTPError(c, 404, message)
}

// MethodNotAllowed response code (405) indicates
// that the request method is known by the server
// but is not supported by the target resource.
func MethodNotAllowed(c echo.Context, message string) (e error) {
	return HTTPError(c, 405, message)
}

// Conflict response code (409) indicates a
// request conflict with current state of the server.
func Conflict(c echo.Context, message string) (e error) {
	return HTTPError(c, 409, message)
}

// Gone response code (410) indicates that access
// to the target resource is no longer available
// at the origin server and that this condition
// is likely to be permanent.
func Gone(c echo.Context, message string) (e error) {
	return HTTPError(c, 410, message)
}

// UnsupportedMediaType response code (415) indicates that
// the server refuses to accept the request because the
// payload format is in an unsupported format.
func UnsupportedMediaType(c echo.Context, message string) (e error) {
	return HTTPError(c, 415, message)
}

// InternalServerError response code (500) indicates that
// the server encountered an unexpected condition that
// prevented it from fulfilling the request.
func InternalServerError(c echo.Context, message string) (e error) {
	return HTTPError(c, 500, message)
}

// NotImplemented response code (501) indicates that the
// server does not support the functionality required
// to fulfill the request.
func NotImplemented(c echo.Context, message string) (e error) {
	return HTTPError(c, 501, message)
}

// HTTPError returns a HTTP error with a descriptive JSON body
func HTTPError(c echo.Context, statusCode int, message string) (e error) {
	resBody := errResponse{
		Message:    message,
		StatusCode: statusCode,
	}

	// Invalid status code supplied, return 500
	if _, ok := codes[statusCode]; !ok {
		return HTTPError(c, 500, "Invalid HTTP status code")
	}

	// Set default error message
	if message == "" {
		resBody.Message = codes[statusCode]
	}

	_, err := json.Marshal(resBody)

	if err != nil {
		return err
	}

	return c.JSON(statusCode, resBody)
}

// GetHTTPResponseFromErr returns the mapped http error to the generic errors model.
func GetHTTPResponseFromErr(c echo.Context, err error) (e error) {

	switch err {
	case models.ErrInternalServerError:
		return InternalServerError(c, err.Error())
	case models.ErrNotFound:
		return NotFound(c, err.Error())
	case models.ErrConflict:
		return Conflict(c, err.Error())
	case models.ErrBadParamInput:
		return BadRequest(c, err.Error())
	case models.ErrUnauthorized:
		return Unauthorized(c, err.Error())
	case models.ErrDatabaseError:
		return InternalServerError(c, err.Error())
	default:
		return InternalServerError(c, "")
	}
}
