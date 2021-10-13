package models

import "time"

type Agreement struct {
	Code        string       `json:"code" bson:"code"`
	Title       *LangOptions `json:"title" bson:"title"`
	Description *LangOptions `json:"description" bson:"description"`
}

type Specification struct {
	ID                   string      `json:"id" bson:"_id,omitempty"`
	Agreement            Agreement   `json:"agreement" bson:"agreement"`
	AdditionalAgreements []Agreement `json:"additional_agreements" bson:"additional_agreements"`
	//VerificationRequired bool        `json:"verification_required" bson:"verification_required"`
	//ValidDaysCount       *int        `json:"valid_days_count" bson:"valid_days_count"`
	UpdatedAt            time.Time   `json:"updated_at" bson:"updated_at"`
	CreatedAt            time.Time   `json:"created_at" bson:"created_at"`
}

type LangOptions struct {
	EN string `json:"en" bson:"en"`
	RU string `json:"ru" bson:"ru"`
	KZ string `json:"kz" bson:"kz"`
}

