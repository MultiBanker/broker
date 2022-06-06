package repository

import (
	"fmt"
	"strings"
)

type ErrInvalidDatastoreName []string

func (ds ErrInvalidDatastoreName) Error() error {
	return fmt.Errorf("datastore: invalid datastore name. Must be one of: %s", strings.Join(ds, ", "))
}

var (
	ErrDatastoreNotImplemented = fmt.Errorf("datastore not implemented")
)
