package utils

import "errors"

var (
	ErrInvalidZipCode    = errors.New("invalid zipcode")
	ErrorNotFoundZipCode = errors.New("can not find zipcode")
)
