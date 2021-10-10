package models

import "time"

type PartnerOrder struct {
	ID          string `json:"-" bson:"_id"`
	ReferenceID string `json:"-" bson:"reference_id"`
	PartnerCode string `json:"-" bson:"partner_code"`
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
	ProductType    string `json:"productType" bson:"product_type"`
	Product        string `json:"product" bson:"product"`
	LoanAmount     string `json:"loanAmount" bson:"loan_amount"`
	LoanLength     string `json:"loanLength" bson:"loan_length"`
	ContractNumber string `json:"contractNumber" bson:"contract_number"`
	MonthlyPayment int    `json:"monthlyPayment" bson:"monthly_payment"`
}
