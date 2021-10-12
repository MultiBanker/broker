package models

import "time"

type Offer struct {
	ID                   string    `json:"id" bson:"_id"`
	PartnerCode          string    `json:"partner_code" bson:"partner_code"`
	Name                 string    `json:"name" bson:"name"`
	PaymentTypeGroupCode string    `json:"payment_type_group_code" bson:"payment_type_group_code"`
	MinOrderSum          int       `json:"min_order_sum" bson:"min_order_sum"`
	MaxOrderSum          int       `json:"max_order_sum" bson:"max_order_sum"`
	CreatedAt            time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" bson:"updated_at"`
}
