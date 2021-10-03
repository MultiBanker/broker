package models

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/dongri/phonenumber"
)

type Partner struct {
	ID          string `json:"id" bson:"_id"`
	CompanyName string `json:"company_name" bson:"company_name"`
	Phone       string `json:"phone" bson:"phone"`
	Email       string `json:"email" bson:"email"`
	BIN         string `json:"bin" bson:"bin"`
	Commission  int    `json:"commission" bson:"commission"`
	LogoURL     string `json:"logo_url" bson:"logo_url"`
	URL         *URL   `json:"url" bson:"url"`
	Enabled     bool   `json:"enabled" bson:"enabled"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type URL struct {
	Create string `json:"create" bson:"create"`
	Update string `json:"update" bson:"update"`
}

func (p Partner) Validate() error {
	if p.CompanyName == "" {
		return ErrCustomValidate("Company Name")
	}
	if p.ValidatePhone(p.Phone) != nil {
		return ErrCustomValidate("phone")
	}
	if p.ValidateEmail(p.Email) != nil {
		return ErrCustomValidate("email")
	}
	if p.BIN == "" {
		return ErrCustomValidate("BIN")
	}
	if p.Commission == 0 {
		return ErrCustomValidate("Commission")
	}
	if p.LogoURL == "" {
		return ErrCustomValidate("LogoURL")
	}
	if p.Username == "" {
		return ErrCustomValidate("Username")
	}
	if p.Password == "" {
		return ErrCustomValidate("Password")
	}
	return nil
}

func (p Partner) ValidatePhone(phone string) error {
	if normPhoneNum(phone) == "" {
		return ErrCustomValidate("phone")
	}

	return nil
}

func (p Partner) ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return ErrCustomValidate("email")
	}

	return nil
}

func (p Partner) ValidateIIN(iin string) error {
	if len(iin) != 12 {
		return ErrCustomValidate("iin. length must be 12")
	}

	if !isDigit(iin) {
		return ErrCustomValidate("iin. must be digit")
	}

	month, err := strconv.Atoi(iin[2:4])
	if err != nil || month <= 0 || month > 12 {
		return ErrCustomValidate("iin. invalid month in iin")
	}

	day, err := strconv.ParseUint(iin[4:6], 10, 64)
	if err != nil || day <= 0 || day > 31 {
		return ErrCustomValidate("iin. invalid day in iin")
	}

	return nil
}

func isDigit(data string) bool {
	if _, err := strconv.Atoi(data); err != nil {
		return false
	}

	return true
}

// normPhoneNum нормализует телефонные номера.
func normPhoneNum(num string) string {
	if len(num) > 10 && num[0] == '8' {
		return phonenumber.Parse(num[1:], "KZ")
	}

	return phonenumber.Parse(num, "KZ")
}
