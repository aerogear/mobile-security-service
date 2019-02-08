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

func BadRequest(c echo.Context, message string) (e error) {
	return HTTPError(c, 400, message)
}

func Unauthorized(c echo.Context, message string) (e error) {
	return HTTPError(c, 401, message)
}

func Forbidden(c echo.Context, message string) (e error) {
	return HTTPError(c, 403, message)
}

func NotFound(c echo.Context, message string) (e error) {
	return HTTPError(c, 404, message)
}

func MethodNotAllowed(c echo.Context, message string) (e error) {
	return HTTPError(c, 405, message)
}

func NotAcceptable(c echo.Context, message string) (e error) {
	return HTTPError(c, 406, message)
}

func Conflict(c echo.Context, message string) (e error) {
	return HTTPError(c, 409, message)
}

func Gone(c echo.Context, message string) (e error) {
	return HTTPError(c, 410, message)
}

func UnsupportedMediaType(c echo.Context, message string) (e error) {
	return HTTPError(c, 404, message)
}

func InternalServerError(c echo.Context, message string) (e error) {
	return HTTPError(c, 500, message)
}

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
		return
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
	default:
		return InternalServerError(c, "")
	}
}
