package dto

import "github.com/MultiBanker/broker/src/models"

type PartnerOrderRequest struct {
	ReferenceId             string          `json:"referenceId"`
	IsDelivery              bool            `json:"isDelivery"`
	Channel                 string          `json:"channel"`
	PaymentMethod           string          `json:"paymentMethod"`
	ProductType             string          `json:"productType"`
	RedirectUrl             string          `json:"redirectUrl"`
	OrderState              string          `json:"orderState"`
	SalesPlace              string          `json:"salesPlace"`
	VerificationSmsCode     string          `json:"verificationSmsCode"`
	VerificationSmsDateTime string          `json:"verificationSmsDateTime"`
	LoanLength              string          `json:"loanLength"`
	Customer                models.Customer `json:"customer"`
	Address                 models.Address  `json:"address"`
	Goods                   []models.Goods  `json:"goods"`
	TotalCost               string          `json:"TotalCost"`
}

type PartnerOrderResponse struct {
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	RedirectUrl string `json:"redirectUrl"`
	RequestUuid string `json:"requestUuid"`
}
