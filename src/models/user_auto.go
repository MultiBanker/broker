package models

import "time"

type UserAuto struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	ApplicationID string `json:"application_id"`
	VIN           string `json:"vin"`

	CreatedAt time.Time `json:"created_at"`
}
