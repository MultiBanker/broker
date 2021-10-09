package models

import "time"

type Market struct {
	ID             string `json:"id" bson:"_id"`
	Title          string `json:"title" bson:"title"`
	LogoURL        string `json:"logo_url" bson:"logo_url"`
	Code           string `json:"code"`
	Location       string `json:"location" bson:"location"`
	WebAddress     string `json:"web_address" bson:"web_address"`
	BIN            string `json:"bin" bson:"bin"`
	UpdateOrderURL string `json:"update_order_url" bson:"update_order_url"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`

	Contact   ContactInfo `json:"contact" bson:"contact"`
	Enabled   bool        `json:"enabled" bson:"enabled"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
}
