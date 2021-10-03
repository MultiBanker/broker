package models

type Order struct {
	ReferenceID string   `json:"referenceId"`
	State       string   `json:"state"`
	StateTitle  string   `json:"stateTitle"`
	Customer    Customer `json:"customer"`
	Offers      []Offers `json:"offers"`
}
type Customer struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}
type Offers struct {
	ProductType    string `json:"productType"`
	Product        string `json:"product"`
	LoanAmount     string `json:"loanAmount"`
	LoanLength     string `json:"loanLength"`
	ContractNumber string `json:"contractNumber"`
	MonthlyPayment int    `json:"monthlyPayment"`
}
