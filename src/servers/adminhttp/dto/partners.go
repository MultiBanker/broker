package dto

import (
	"fmt"
	"strings"
)

type PartnerRequest struct {
	// Код партнера
	Code        string      `json:"code" bson:"code" example:"mfo_airba"`
	// Название компании
	CompanyName string      `json:"company_name" bson:"company_name" example:"Airba Pay"`
	// Коммиссия
	Commission  int         `json:"commission" bson:"commission" example:"5"`
	// Веб адрес компании
	LogoURL     string      `json:"logo_url" bson:"logo_url" example:"https://emotionsgroup.kz/uploads/7f47072c77ad20f58acb9e7114cfeb5c.jpg"`
	// Адреса на обновления АПИ
	URL         *URL        `json:"url" bson:"url"`
	// Контактные данные
	Contact     ContactInfo `json:"contact" bson:"contact"`
	// Включен
	Enabled     bool        `json:"enabled" bson:"enabled" example:"true"`
	// Имя пользователя
	Username string `json:"username" bson:"username" example:"SuperRinat"`
	// Пароль
	Password string `json:"password" bson:"password" example:"Password123"`
}

type URL struct {
	// урл авторизации
	Auth string `json:"auth" bson:"auth" example:"https://"`
	// Урл который создает заказ на стороне партнера
	Create string `json:"create" bson:"create" example:"https://"`
	// Урл который обновляет заказ на стороне партнера
	Update string `json:"update" bson:"update" example:"https://"`
}

func (u URL) Validate() error {
	var errstrings []string

	if u.Auth == "" {
		errstrings = append(errstrings, ValidationIsEmpty("auth url").Error())
	}

	if u.Create == "" {
		errstrings = append(errstrings, ValidationIsEmpty("create url").Error())
	}

	if u.Update == "" {
		errstrings = append(errstrings, ValidationIsEmpty("update url").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

func (p PartnerRequest) Validate() error {
	var errstrings []string

	if p.Code == "" {
		errstrings = append(errstrings, ValidationIsEmpty("partner code").Error())
	}
	if p.CompanyName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("company name").Error())
	}
	if p.Commission == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("commission").Error())
	}
	if p.LogoURL == "" {
		errstrings = append(errstrings, ValidationIsEmpty("logo url").Error())
	}
	if p.Username == "" {
		errstrings = append(errstrings, ValidationIsEmpty("username").Error())
	}
	if p.Password == "" {
		errstrings = append(errstrings, ValidationIsEmpty("password").Error())
	}
	if p.URL.Validate() != nil {
		errstrings = append(errstrings, p.URL.Validate().Error())
	}
	if p.Contact.Validate() != nil {
		errstrings = append(errstrings, p.Contact.Validate().Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	return nil
}
