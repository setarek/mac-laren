package error

import "errors"

var (
	ErrServer           = errors.New("server error")
	ErrEmptyBodyRequest = errors.New("empty body request")
)
