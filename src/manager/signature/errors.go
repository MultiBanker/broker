package signature

import "errors"

var (
	ErrTryLimitReached      = errors.New("try limit reached")
	ErrTooManyVerifications = errors.New("too many verification requests")
	ErrTokenExpired         = errors.New("token lifetime expired")
	ErrTokenMismatch        = errors.New("token does not match")
	ErrCannotBeSigned       = errors.New("cannot be signed") // incorrect status
)
