package models

import (
	"errors"
	"fmt"
)

var (
	ErrValidator = errors.New("invalid field")
)

func ErrCustomValidate(text string) error {
	return fmt.Errorf("%w: %s", ErrValidator, text)
}


var (
	ErrInvalidToken        = errors.New("invalid token")
	ErrInvalidTime         = errors.New("invalid time in offer")
	ErrInvalidRowsExcel    = errors.New("excel must contain at least one header row and one product row")
	ErrInvalidHeaderExcel  = errors.New("excel header has too few columns")
	ErrInvalidRole         = errors.New("invalid role")
)

