package dto

import (
	"time"

	"github.com/MultiBanker/broker/src/models"
)

type OrderRequest struct {
	ID                      string   `json:"id" bson:"_id"`
	ReferenceID             string   `json:"referenceId" bson:"reference_id"`
	OrderState              string   `json:"orderState" bson:"order_state"`
	RedirectURL             string   `json:"redirectUrl" bson:"redirect_url"`
	Channel                 string   `json:"channel" bson:"channel"`
	ProductType             string   `json:"productType" bson:"product_type"`
	PaymentMethod           string   `json:"paymentMethod" bson:"payment_method"`
	IsDelivery              bool     `json:"isDelivery" bson:"is_delivery"`
	TotalCost               string   `json:"totalCost" bson:"total_cost"`
	LoanLength              string   `json:"loanLength" bson:"loan_length"`
	SalesPlace              string   `json:"salesPlace" bson:"sales_place"`
	VerificationSMSCode     string   `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string   `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                Customer `json:"customer" bson:"customer"`
	Address                 Address  `json:"address" bson:"address"`
	Goods                   []Goods  `json:"goods" bson:"goods"`

	BankType  string    `json:"-" bson:"bank_type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Contact struct {
	MobileNumber string `json:"mobileNumber" bson:"mobile_number"`
	Email        string `json:"email" bson:"email"`
}
type Customer struct {
	TaxCode    string  `json:"taxCode" bson:"tax_code"`
	Firstname  string  `json:"firstName" bson:"firstname"`
	Lastname   string  `json:"lastName" bson:"lastname"`
	MiddleName string  `json:"middleName" bson:"middle_name"`
	State      string  `json:"state" bson:"state"`
	Contact    Contact `json:"contact" bson:"contact"`
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

type OrderResponse struct {
	Status      string `json:"status"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirectUrl"`
	RequestUUID string `json:"requestUuid"`
	Message     string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type IDResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Markets struct {
	Total   int64           `json:"total"`
	Markets []models.Market `json:"markets"`
}

type Orders struct {
	Total  int64           `json:"total"`
	Orders []*OrderRequest `json:"orders"`
}

type Partners struct {
	Total    int64            `json:"total"`
	Partners []models.Partner `json:"partners"`
}
