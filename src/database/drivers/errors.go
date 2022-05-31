package drivers

import "errors"

var (
	ErrDoesNotExist  = errors.New("the does not exist")
	ErrAlreadyExists = errors.New("already exists")
	ErrBadID         = errors.New("err bad id")
)
