package exception

import "errors"

var ErrInternalServer = errors.New("internal server error")

var ErrIdIsNotFound = errors.New("id is not found")

var ErrBadRequest = errors.New("operation falied, request resource not valid")
