package models

import "time"

type Order struct {
	ID                      string            `json:"id" bson:"_id"`
	ReferenceID             string            `json:"referenceId" bson:"reference_id"`
	OrderState              string            `json:"orderState" bson:"order_state"`
	RedirectURL             string            `json:"redirectUrl" bson:"redirect_url"`
	Channel                 string            `json:"channel" bson:"channel"`
	SystemCode              string            `json:"system_code"`
	StateCode               string            `json:"-"`
	MarketCode              string            `json:"-"`
	ProductType             string            `json:"productType" bson:"product_type"`
	PaymentMethod           string            `json:"paymentMethod" bson:"payment_method"`
	IsDelivery              bool              `json:"isDelivery" bson:"is_delivery"`
	TotalCost               string            `json:"totalCost" bson:"total_cost"`
	LoanLength              string            `json:"loanLength" bson:"loan_length"`
	SalesPlace              string            `json:"salesPlace" bson:"sales_place"`
	VerificationId          string            `json:"verificationId"`
	VerificationSMSCode     string            `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string            `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                Customer          `json:"customer" bson:"customer"`
	Address                 Address           `json:"address" bson:"address"`
	Goods                   []Goods           `json:"goods" bson:"goods"`
	PaymentPartners         []PaymentPartners `json:"paymentPartners"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Customer struct {
	IIN        string  `json:"iin"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName string  `json:"middleName"`
	Contact    Contact `json:"contact"`
}

type Address struct {
	Delivery    string `json:"delivery" bson:"delivery"`
	PickupPoint string `json:"pickupPoint" bson:"pickup_point"`
}
type Goods struct {
	Category string `json:"category" bson:"category"`
	Brand    string `json:"brand" bson:"brand"`
	Price    string `json:"price" bson:"price"`
	Model    string `json:"model" bson:"model"`
	Image    string `json:"image" bson:"image"`
}

type Contact struct {
	MobileNumber string `json:"mobileNumber" bson:"mobile_number"`
	Email        string `json:"email" bson:"email"`
}

type PaymentPartners struct {
	Code string `json:"code"`
}
