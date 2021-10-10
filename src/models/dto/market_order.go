package dto

import "github.com/MultiBanker/broker/src/models"

type MarketOrderRequest struct {
	Amount                  string                   `json:"amount"`
	IsDelivery              bool                     `json:"isDelivery"`
	CityID                  string                   `json:"cityId"`
	Channel                 string                   `json:"channel"`
	PaymentMethod           string                   `json:"paymentMethod"`
	ProductType             string                   `json:"productType"`
	RedirectURL             string                   `json:"redirectUrl"`
	SystemCode              string                   `json:"systemCode"`
	VerificationSmsCode     string                   `json:"verificationSmsCode"`
	VerificationID          string                   `json:"verificationId"`
	LoanLength              int                      `json:"loanLength"`
	VerificationSmsDateTime string                   `json:"verificationSmsDateTime"`
	Customer                models.Customer          `json:"customer"`
	Address                 models.Address           `json:"address"`
	Goods                   []models.Goods           `json:"goods"`
	PaymentPartners         []models.PaymentPartners `json:"paymentPartners"`
}

type MarketOrderResponse struct {
	ReferenceID string `json:"reference_id"`
}
