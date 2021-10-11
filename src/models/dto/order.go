package dto

import (
	"time"

	"github.com/MultiBanker/broker/src/models"
)

type OrderBankRequest struct {
	ID                      string   `json:"id" bson:"_id"`
	ReferenceID             string   `json:"referenceId" bson:"reference_id"`
	OrderState              string   `json:"orderState" bson:"order_state"`
	RedirectURL             string   `json:"redirectUrl" bson:"redirect_url"`
	Channel                 string   `json:"channel" bson:"channel"`
	ProductType             string   `json:"productType" bson:"product_type"`
	PaymentMethod           string   `json:"paymentMethod" bson:"payment_method"`
	IsDelivery              bool     `json:"isDelivery" bson:"is_delivery"`
	TotalCost               string   `json:"totalCost" bson:"total_cost"`
	LoanLength              int      `json:"loanLength" bson:"loan_length"`
	SalesPlace              string   `json:"salesPlace" bson:"sales_place"`
	VerificationSMSCode     string   `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string   `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                Customer `json:"customer" bson:"customer"`
	Address                 Address  `json:"address" bson:"address"`
	Goods                   []Goods  `json:"goods" bson:"goods"`
}

type OrderRequest struct {
	ID                      string            `json:"id" bson:"_id"`
	SystemCode              string            `json:"systemCode"`
	Channel                 string            `json:"channel"`
	StateCode               string            `json:"-"`
	RedirectURL             string            `json:"redirectUrl"`
	IsDelivery              bool              `json:"isDelivery"`
	ProductType             string            `json:"productType"`
	PaymentMethod           string            `json:"paymentMethod"`
	OrderID                 string            `json:"order_id"`
	Amount                  string            `json:"amount"`
	VerificationSmsCode     string            `json:"verificationSmsCode"`
	VerificationSmsDateTime string            `json:"verificationSmsDateTime"`
	Customer                Customer          `json:"customer"`
	Address                 Address           `json:"address"`
	Goods                   []Goods           `json:"goods"`
	PaymentPartners         []PaymentPartners `json:"paymentPartners"`

	CreatedAt time.Time `json:"-" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}

func (o OrderRequest) ToBankOrder() OrderBankRequest {
	return OrderBankRequest{
		ReferenceID: o.ID,
	}
}

type Contact struct {
	MobileNumber string `json:"mobileNumber" bson:"mobile_number"`
	Email        string `json:"email" bson:"email"`
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

type Customer struct {
	IIN        string  `json:"iin"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName string  `json:"middleName"`
	Contact    Contact `json:"contact"`
}

type PaymentPartners struct {
	Code string `json:"code"`
}

type OrderResponse struct {
	ID          string `json:"-" bson:"_id"`
	ReferenceID string `json:"-" bson:"reference_id"`
	PartnerCode string `json:"-" bson:"partner_code"`
	MarketCode  string `json:"-" bson:"market_code"`

	Status      string   `json:"status" bson:"status"`
	Code        string   `json:"code" bson:"code"`
	RedirectURL string   `json:"redirectUrl" bson:"redirect_url"`
	RequestUUID string   `json:"requestUuid" bson:"request_uuid"`
	State       string   `json:"state" bson:"state"`
	StateTitle  string   `json:"state_title"`
	Message     string   `json:"message" bson:"message"`
	Offers      []Offers `json:"offers" bson:"offers"`

	CreatedAt time.Time `json:"-" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
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

type Orders struct {
	Total  int64           `json:"total"`
	Orders []*OrderRequest `json:"orders"`
}

type Partners struct {
	Total    int64            `json:"total"`
	Partners []models.Partner `json:"partners"`
}
