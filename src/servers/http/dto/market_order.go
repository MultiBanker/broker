package dto

import (
	"fmt"
	"strings"

	"github.com/MultiBanker/broker/src/models"
)

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

func (m MarketOrderRequest) Validate() error {
	var errstrings []string

	if m.Amount == "" {
		errstrings = append(errstrings, ValidationIsEmpty("amount").Error())
	}

	if m.CityID == "" {
		errstrings = append(errstrings, ValidationIsEmpty("city id").Error())
	}

	if m.Channel == "" {
		errstrings = append(errstrings, ValidationIsEmpty("channel").Error())
	}

	if m.PaymentMethod == "" {
		errstrings = append(errstrings, ValidationIsEmpty("payment method").Error())
	}

	if m.ProductType == "" {
		errstrings = append(errstrings, ValidationIsEmpty("product type").Error())
	}

	if m.RedirectURL == "" {
		errstrings = append(errstrings, ValidationIsEmpty("redirect url").Error())
	}

	if m.SystemCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("system code").Error())
	}

	if m.VerificationSmsCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("verification sms code").Error())
	}

	if m.VerificationID == "" {
		errstrings = append(errstrings, ValidationIsEmpty("verification id").Error())
	}

	if m.VerificationSmsDateTime == "" {
		errstrings = append(errstrings, ValidationIsEmpty("verification date time").Error())
	}

	if m.LoanLength == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("verification id").Error())
	}

	if err := m.Customer.Validate(); err != nil {
		errstrings = append(errstrings, m.Customer.Validate().Error())
	}

	if err := m.Address.Validate(); err != nil {
		errstrings = append(errstrings, m.Address.Validate().Error())
	}

	if m.Goods == nil {
		errstrings = append(errstrings, ValidationIsEmpty("goods").Error())
	}

	if m.Goods != nil {
		for _, good := range m.Goods {
			errstrings = append(errstrings, good.Validate().Error())
		}
	}

	if m.PaymentPartners == nil {
		errstrings = append(errstrings, ValidationIsEmpty("payment partners").Error())
	}

	if m.PaymentPartners != nil {
		for _, partner := range m.PaymentPartners {
			errstrings = append(errstrings, partner.Validate().Error())
		}
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

type MarketOrderResponse struct {
	ReferenceID string `json:"reference_id"`
}
