package models

import "errors"

var (
	// ErrInternalServerError returns a new Internal Server Error
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound returns a new Not Found Error
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict returns a new Conflict error
	ErrConflict = errors.New("Your Item already exists")
	// ErrBadParamInput returns a Bad Parameter Input Error
	ErrBadParamInput = errors.New("Given Param is not valid")
	// ErrUnauthorized returns a new Unauthorized Error
	ErrUnauthorized = errors.New("Missing or Invalid authentication token")
	// ErrDatabaseError returns a New Database Error
	ErrDatabaseError = errors.New("An error has occurred in the database")
)
