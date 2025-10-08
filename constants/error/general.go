package error

import "errors"

const (
	InternalServerError = "internal_server_error"
	BadRequest          = "bad_request"
	Unauthorized        = "unauthorized"
	NotFound            = "not_found"
	Conflict            = "conflict"
	Success             = "success"
)

var (
	ErrServerInternalError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrSuccess             = errors.New("success")
)

var GeneralError = []error{
	ErrServerInternalError,
	ErrBadRequest,
	ErrUnauthorized,
	ErrNotFound,
	ErrConflict,
	ErrSuccess,
}
