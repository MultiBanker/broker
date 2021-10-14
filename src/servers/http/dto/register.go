package dto

import (
	"fmt"
	"strings"
)

type SignUpPartnerRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	StoreName  string `json:"store_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	BIN        string `json:"bin,omitempty"`
	Commission int    `json:"commission"`
	LogoURL    string `json:"logo_url"`
}

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
