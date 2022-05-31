package database

import (
	"fmt"
	"strings"
)

type ErrInvalidDatastoreName []string

func (ds ErrInvalidDatastoreName) Error() error {
	return fmt.Errorf("datastore: invalid datastore name. Must be one of: %s", strings.Join(ds, ", "))
}

var (
	ErrEmptyConfigStruct       = fmt.Errorf("empty config structure")
	ErrDatastoreNotImplemented = fmt.Errorf("datastore not implemented")
	ErrDoesNotExist        = fmt.Errorf("user not exists")
)
