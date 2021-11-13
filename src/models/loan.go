package models

import "time"

type LoanProgram struct {
	ID          string    `json:"id" bson:"id"`
	Code        string    `json:"code" bson:"code"`
	IsEnabled   bool      `json:"is_enabled" bson:"is_enabled"`
	MaxAmount   int       `json:"max_amount" bson:"max_amount"`
	MinAmount   int       `json:"min_amount" bson:"min_amount"`
	Name        string    `json:"name" bson:"name"`
	Note        string    `json:"note" bson:"note"`
	PartnerCode string    `json:"partner_code" bson:"partner_code"`
	Rate        float64   `json:"rate" bson:"rate"`
	Term        int       `json:"term" bson:"term"`
	Type        string    `json:"type" bson:"type"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
