package models

import "time"

type Recovery struct {
	ID            string    `json:"id" bson:"_id"`                          // ID
	UserID        string    `json:"user_id" bson:"user_id"`                 // ID пользователя
	Destination   string    `json:"destination" bson:"destination"`         // Адресат отправки ОТП
	Channel       string    `json:"channel" bson:"channel"`                 // Канал отправки ОТП
	OTP           string    `json:"otp" bson:"otp"`                         // Сгенерированный ОТП
	Status        string    `json:"status" bson:"status"`                   // Статус обработки верификации
	Send          bool      `json:"send" bson:"send"`                       // Флаг отправки ОТП
	Tries         uint8     `json:"tries" bson:"tries"`                     // Количество попыток для ввода смс
	Count         uint8     `json:"count" bson:"count"`                     // Количество сгенерированных ОТП
	ExpiredAt     time.Time `json:"expired_at" bson:"expired_at"`           // Время истечения
	NextAttemptAt time.Time `json:"next_attempt_at" bson:"next_attempt_at"` // Время последующей попытки
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`           // Дата создания
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`           // Дата обновления
}
