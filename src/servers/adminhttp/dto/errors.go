package dto

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrIsEmpty = fmt.Errorf("is Empty")

func ValidationIsEmpty(value string) error {
	return errors.Wrap(ErrIsEmpty, value)
}
