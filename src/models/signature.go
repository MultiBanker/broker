package models

import "time"

type Signature struct {
	ID              string        `json:"id" bson:"_id,omitempty"`
	VerificationID  string        `json:"verification_id" bson:"verification_id"`
	Verification    Verification  `json:"verification" bson:"verification"`
	Agree           Specification `json:"agree" bson:"agree"`
	UserID          string        `json:"user_id" bson:"user_id"`
	AdditionalCodes []string      `json:"additional_codes" bson:"additional_codes"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`
}

type Verification struct {
	ID             string    `json:"id" bson:"_id"`
	UserID         string    `json:"user_id" bson:"user_id"`
	Phone          string    `json:"phone" bson:"phone"`
	Token          string    `json:"token" bson:"token"`
	Status         string    `json:"status" bson:"status"`
	Tries          int       `json:"tries" bson:"tries"`
	NotificationID string    `json:"notification_id" bson:"notification_id"`
	NextAttemptAt  time.Time `json:"next_attempt_at" bson:"next_attempt_at"`
}
