package models

import (
	"fmt"
	"strings"
	"time"
)

type PartnerOrder struct {
	ID          string `json:"-" bson:"_id"`
	ReferenceID string `json:"reference_id" bson:"reference_id"`
	PartnerCode string `json:"partner_code" bson:"partner_code"`
	MarketCode  string `json:"-" bson:"market_code"`

	Status      string   `json:"status" bson:"status"`
	Code        int      `json:"code" bson:"code"`
	RedirectURL string   `json:"redirectUrl" bson:"redirect_url"`
	RequestUUID string   `json:"requestUuid" bson:"request_uuid"`
	State       string   `json:"state" bson:"state"`
	StateTitle  string   `json:"state_title"`
	Message     string   `json:"message" bson:"message"`
	Offers      []Offers `json:"offers" bson:"offers"`
	Customer    Customer `json:"customer" bson:"customer"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Offers struct {
	ProductType    string `json:"productType" bson:"product_type" example:"installment"`
	Product        string `json:"product" bson:"product" example:"rassrochka_12"`
	LoanAmount     string `json:"loanAmount" bson:"loan_amount" example:"144000"`
	LoanLength     string `json:"loanLength" bson:"loan_length" example:"12"`
	ContractNumber string `json:"contractNumber" bson:"contract_number" example:"d12ed1"`
	MonthlyPayment int    `json:"monthlyPayment" bson:"monthly_payment" example:"12000"`
}

func (o Offers) Validate() error {
	var errstrings []string

	_, err := ValidateProductType(o.ProductType)
	if err != nil {
		errstrings = append(errstrings, err.Error())
	}

	if o.Product == "" {
		errstrings = append(errstrings, ValidationIsEmpty("product").Error())
	}


	if o.LoanAmount == "" {
		errstrings = append(errstrings, ValidationIsEmpty("loan amount").Error())
	}

	if o.LoanLength == "" {
		errstrings = append(errstrings, ValidationIsEmpty("loan length").Error())
	}


	if o.ContractNumber == "" {
		errstrings = append(errstrings, ValidationIsEmpty("contract number").Error())
	}

	if o.MonthlyPayment == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("monthly payment").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
