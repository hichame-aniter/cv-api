package errors

import "errors"

var (
	ErrNotFound = errors.New("Resource not found")
	ErrInvalid  = errors.New("Invalid request")
)
