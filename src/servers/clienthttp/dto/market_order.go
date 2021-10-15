package dto

import (
	"fmt"
	"strings"

	"github.com/MultiBanker/broker/src/models"
)

type MarketOrderRequest struct {
	Amount                  string                   `json:"amount" example:"5000"`
	IsDelivery              bool                     `json:"isDelivery" example:"true"`
	CityID                  string                   `json:"cityId" example:"050000"`
	Channel                 string                   `json:"channel" example:"airba_web"`
	PaymentMethod           string                   `json:"paymentMethod" example:"annuity"`
	ProductType             string                   `json:"productType" example:"installment"`
	RedirectURL             string                   `json:"redirectUrl" example:"https://airba.kz/order/ok"`
	SystemCode              string                   `json:"systemCode" example:"oms"`
	VerificationSmsCode     string                   `json:"verificationSmsCode" example:"12321"`
	VerificationID          string                   `json:"verificationId" example:"dsad12"`
	LoanLength              int                      `json:"loanLength" example:"12"`
	VerificationSmsDateTime string                   `json:"verificationSmsDateTime" example:"12.12.2020"`
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
