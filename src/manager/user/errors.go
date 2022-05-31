package user

import "errors"

var (
	ErrUserDisabled           = errors.New("user is disabled")
	ErrUserPhoneIsNotVerified = errors.New("user's phone is not verified")
	ErrUserEmailIsNotVerified = errors.New("user's email is not verified")
	ErrInvalidLoginOrPassword = errors.New("login or password is incorrect")
	ErrInvalidOTP             = errors.New("provided otp is invalid")
	ErrExpiredOTP             = errors.New("provided otp is expired")
	ErrTriesLimitIsOver       = errors.New("tries limit is over, try this action after 60 min")
	ErrUnknownCompanyID       = errors.New("unknown company ID")
	ErrAuthorization          = errors.New("wrong username and password")
)

type TooManyRequestsError struct {
	NextAttemptAt string `json:"next_attempt_at,omitempty"`
}

func (e TooManyRequestsError) Error() string {
	return e.NextAttemptAt
}

