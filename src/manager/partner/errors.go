package partner

import "errors"

var (
	ErrAuthorization          = errors.New("wrong username and password")
)