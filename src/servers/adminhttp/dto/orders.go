package dto

import "github.com/MultiBanker/broker/src/models"

type OrderAdminUpdateRequest struct {
	ReferenceID             string                   `json:"referenceId" bson:"reference_id"`
	OrderState              string                   `json:"orderState" bson:"order_state"`
	RedirectURL             string                   `json:"redirectUrl" bson:"redirect_url"`
	Channel                 string                   `json:"channel" bson:"channel"`
	SystemCode              string                   `json:"systemCode"`
	StateCode               string                   `json:"-"`
	MarketCode              string                   `json:"-"`
	ProductType             string                   `json:"productType" bson:"product_type"`
	PaymentMethod           string                   `json:"paymentMethod" bson:"payment_method"`
	IsDelivery              bool                     `json:"isDelivery" bson:"is_delivery"`
	TotalCost               string                   `json:"totalCost" bson:"total_cost"`
	LoanLength              string                   `json:"loanLength" bson:"loan_length"`
	SalesPlace              string                   `json:"salesPlace" bson:"sales_place"`
	VerificationId          string                   `json:"verificationId"`
	VerificationSMSCode     string                   `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string                   `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                models.Customer          `json:"customer" bson:"customer"`
	Address                 models.Address           `json:"address" bson:"address"`
	Goods                   []models.Goods           `json:"goods" bson:"goods"`
	PaymentPartners         []models.PaymentPartners `json:"paymentPartners"`
}
