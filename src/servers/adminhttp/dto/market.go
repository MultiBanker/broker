package dto

import (
	"fmt"
	"strings"
)

type MarketRequest struct {
	// Название компании
	CompanyName string `json:"company_name" bson:"company_name" example:"Tinka"`
	// Веб адрес картинки
	LogoURL string `json:"logo_url" bson:"logo_url" example:"https://plusworld.ru/wp-content/uploads/2021/09/759950f2eafdab128b887f1296316ea6f1af5152.jpg"`
	// Код картинки
	Code string `json:"code" example:"tinkoff"`
	// Обновление заказов
	UpdateOrderURL string `json:"update_order_url" bson:"update_order_url" example:"https://google.com"`
	// Имя пользователя
	Username string `json:"username" bson:"username" example:"Olejka"`
	// Пароль
	Password string `json:"password" bson:"password" example:"Password123"`
	// Контакт
	Contact ContactInfo `json:"contact" bson:"contact"`
	// Включение
	Enabled bool `json:"enabled" bson:"enabled" example:"true"`
}

func (m MarketRequest) Validate() error {
	var errstrings []string

	if m.CompanyName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("company name").Error())
	}
	if m.LogoURL == "" {
		errstrings = append(errstrings, ValidationIsEmpty("logo url").Error())
	}
	if m.Code == "" {
		errstrings = append(errstrings, ValidationIsEmpty("market code").Error())
	}
	if m.UpdateOrderURL == "" {
		errstrings = append(errstrings, ValidationIsEmpty("update order url").Error())
	}
	if m.Username == "" {
		errstrings = append(errstrings, ValidationIsEmpty("username").Error())
	}
	if m.Password == "" {
		errstrings = append(errstrings, ValidationIsEmpty("password").Error())
	}
	if m.Contact.Validate() != nil {
		errstrings = append(errstrings, m.Contact.Validate().Error())
	}
	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
