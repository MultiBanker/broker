package dto

import (
	"fmt"
	"strings"
)

type ContactInfo struct {
	// Имя
	FirstName string `json:"first_name" bson:"first_name" example:"Oleg"`
	// Фамилия
	LastName string `json:"last_name" bson:"last_name" example:"Tinkoff"`
	// Имейл
	Email string `json:"email" bson:"email" example:"oleg@tinkoff.ru"`
	// Телефон
	Phone string `json:"phone" bson:"phone" example:"87777777777"`
	// Адрес организации
	Location string `json:"location" bson:"location" example:"Bldg. 26, 38A, 2 Khutorskaya str., Moscow, Russia."`
	// Веб адрес организации
	WebAddress string `json:"web_address" bson:"web_address" example:"https://www.tinkoff.ru"`
	// БИН организации
	BIN string `json:"bin" bson:"bin" example:"5189011425"`
}

func (c ContactInfo) Validate() error {
	var errstrings []string

	if c.FirstName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("first name").Error())
	}
	if c.LastName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("last name").Error())
	}
	if c.Email == "" {
		errstrings = append(errstrings, ValidationIsEmpty("email").Error())
	}

	if c.Phone == "" {
		errstrings = append(errstrings, ValidationIsEmpty("phone").Error())
	}

	if c.WebAddress == "" {
		errstrings = append(errstrings, ValidationIsEmpty("web address").Error())
	}

	if c.Location == "" {
		errstrings = append(errstrings, ValidationIsEmpty("location").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
