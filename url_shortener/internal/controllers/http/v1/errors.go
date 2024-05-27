package http_v1

import "errors"

var (
	ErrInternalError  = errors.New("internal error")
	ErrNotValidURL    = errors.New("not valid URL")
	ErrInvalidRequest = errors.New("invalid request")
	ErrRequiredUrl    = errors.New("url is required")
	ErrNotFound       = errors.New("not found")
)
