package dto

import (
	"fmt"
	"strings"
)

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func (l Login) Validate() error {
	var errstrings []string

	if l.Username == "" {
		errstrings = append(errstrings, ValidationIsEmpty("username").Error())
	}

	if l.Password == "" {
		errstrings = append(errstrings, ValidationIsEmpty("password").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
