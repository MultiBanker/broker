package dto

import "github.com/MultiBanker/broker/src/models"

type OrderPartnerUpdateRequest struct {
	ReferenceID string          `json:"referenceId" bson:"reference_id"`
	State       string          `json:"state" bson:"state"`
	StateTitle  string          `json:"stateTitle" bson:"state_title"`
	Customer    FIO             `json:"customer" bson:"customer"`
	Offers      []models.Offers `json:"offers" bson:"offers"`
}
type FIO struct {
	FirstName  string `json:"firstName" bson:"first_name"`
	LastName   string `json:"lastName" bson:"last_name"`
	MiddleName string `json:"middleName" bson:"middle_name"`
}
type Offers struct {
	ProductType    string `json:"productType" bson:"product_type"`
	Product        string `json:"product" bson:"product"`
	LoanAmount     string `json:"loanAmount" bson:"loan_amount"`
	LoanLength     string `json:"loanLength" bson:"loan_length"`
	ContractNumber string `json:"contractNumber" bson:"contract_number"`
	MonthlyPayment int    `json:"monthlyPayment" bson:"monthly_payment"`
}

type UpdateMarketOrderRequest struct {
	ReferenceId string `json:"referenceId"`
	State       string `json:"state"`
	StateTitle  string `json:"stateTitle"`
	ProductCode string `json:"productCode"`
	LoanLength  string `json:"loanLength"`
	Reason      string `json:"reason"`
}
