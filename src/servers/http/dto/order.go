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
	LoanLength              string   `json:"loanLength" bson:"loan_length"`
	SalesPlace              string   `json:"salesPlace" bson:"sales_place"`
	VerificationSMSCode     string   `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string   `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                Customer `json:"customer" bson:"customer"`
	Address                 Address  `json:"address" bson:"address"`
	Goods                   []Goods  `json:"goods" bson:"goods"`
}

func (o OrderRequest) ToBankOrder() OrderBankRequest {
	return OrderBankRequest{}
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

type OrderRequest struct {
	Amount                  string                   `json:"amount"`
	IsDelivery              bool                     `json:"isDelivery"`
	CityId                  string                   `json:"cityId"`
	Channel                 string                   `json:"channel"`
	PaymentMethod           string                   `json:"paymentMethod"`
	ProductType             string                   `json:"productType"`
	RedirectUrl             string                   `json:"redirectUrl"`
	SystemCode              string                   `json:"systemCode"`
	VerificationSmsCode     string                   `json:"verificationSmsCode"`
	VerificationId          string                   `json:"verificationId"`
	LoanLength              int                      `json:"loanLength"`
	VerificationSmsDateTime string                   `json:"verificationSmsDateTime"`
	Customer                models.Customer          `json:"customer"`
	Address                 models.Address           `json:"address"`
	Goods                   []models.Goods           `json:"goods"`
	PaymentPartners         []models.PaymentPartners `json:"paymentPartners"`
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
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Price    int    `json:"price"`
	Model    string `json:"model"`
	Image    string `json:"image"`
	Sku      string `json:"sku"`
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
