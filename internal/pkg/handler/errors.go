package handler

import "errors"

var (
	errInvalidInputBody    = errors.New("invalid input body")
	errEmptyAuthHeader     = errors.New("auth header is empty")
	errInvalidAuthHeader   = errors.New("invalid auth header")
	errIdNotFoundInContext = errors.New("id not found in context")
	errIdType              = errors.New("incorrect id: id must be integer number")
)
