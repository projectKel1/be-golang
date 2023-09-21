package exception

import "errors"

var (
	ErrInternalServer  = errors.New("internal server error")
	ErrIdIsNotFound    = errors.New("id is not found")
	ErrBadRequest      = errors.New("operation falied, request resource not valid")
	ErrForbiddenAccess = errors.New("access denied!")
)
