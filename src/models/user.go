package models

import "time"

type User struct {
	ID              string    `json:"id" bson:"_id"`
	FirstName       string    `json:"first_name" bson:"first_name"`               // Имя пользователя
	LastName        string    `json:"last_name" bson:"last_name"`                 // Фамилия пользователя
	Patronymic      string    `json:"patronymic" bson:"patronymic"`               // Отчество пользователя
	IIN             string    `json:"iin" bson:"iin"`                             // ИИН
	Phone           string    `json:"phone" bson:"phone"`                         // Номер телефона пользователя
	Email           string    `json:"email" bson:"email"`                         // Почта пользователя
	Password        string    `json:"password" bson:"password"`                   // Пароль пользователя
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`               // Дата создания пользователя
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`               // Дата обновления данных пользователя
	IsEnabled       bool      `json:"is_enabled" bson:"is_enabled"`               // Статус пользователя(включен/отключен)
	IsPhoneVerified bool      `json:"is_phone_verified" bson:"is_phone_verified"` // Статус верификации номера телефона
	IsEmailVerified bool      `json:"is_email_verified" bson:"is_email_verified"` // Статус верификации электронной почты
}

type UserInfo struct {
	ID              string    `json:"id"`
	FirstName       string    `json:"first_name"`        // Имя пользователя
	LastName        string    `json:"last_name"`         // Фамилия пользователя
	Patronymic      string    `json:"patronymic"`        // Отчество пользователя
	Phone           string    `json:"phone"`             // Номер телефона пользователя
	Email           string    `json:"email"`             // Почта пользователя
	CreatedAt       time.Time `json:"created_at"`        // Дата создания пользователя
	UpdatedAt       time.Time `json:"updated_at"`        // Дата обновления данных пользователя
	Language        string    `json:"language"`          // Язык интерфейса
	IsEnabled       bool      `json:"is_enabled"`        // Статус пользователя(включен/отключен)
	IsPhoneVerified bool      `json:"is_phone_verified"` // Статус верификации номера телефона
	IsEmailVerified bool      `json:"is_email_verified"` // Статус верификации электронной почты
}

type UserApplication struct {
	UserID    string
	ChosenSKU string
}
