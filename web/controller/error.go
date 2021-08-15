package controller

import "errors"

// list of errors
var (
	ErrInvalidData    = errors.New("invalid data")
	ErrNotFound       = errors.New("not found")
	ErrForbidden      = errors.New("forbidden")
	ErrProcessFailure = errors.New("failed to process")
)
