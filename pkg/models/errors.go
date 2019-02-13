package models

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exists")
	ErrBadParamInput       = errors.New("Given Param is not valid")
	ErrUnauthorized        = errors.New("Missing or Invalid authentication token")
	ErrDatabaseError       = errors.New("An error has occurred in the database")
)
