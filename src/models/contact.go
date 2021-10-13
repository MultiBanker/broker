package models

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
