package models

import "time"

type Market struct {
	ID             string      `json:"id" bson:"_id"`
	CompanyName    string      `json:"title" bson:"title"`
	LogoURL        string      `json:"logo_url" bson:"logo_url"`
	Code           string      `json:"code"`
	UpdateOrderURL string      `json:"update_order_url" bson:"update_order_url"`
	Username       string      `json:"username" bson:"username"`
	Password       string      `json:"password" bson:"password"`
	Contact        ContactInfo `json:"contact" bson:"contact"`
	Enabled        bool        `json:"enabled" bson:"enabled"`
	CreatedAt      time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at" bson:"updated_at"`
}
