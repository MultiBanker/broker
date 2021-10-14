package dto

import (
	"fmt"
	"strings"

	"github.com/MultiBanker/broker/src/models"
)

type OrderPartnerUpdateRequest struct {
	ReferenceID string          `json:"referenceId" bson:"reference_id"`
	State       string          `json:"state" bson:"state"`
	StateTitle  string          `json:"stateTitle" bson:"state_title"`
	Customer    FIO             `json:"customer" bson:"customer"`
	Offers      []models.Offers `json:"offers" bson:"offers"`
}

func (o OrderPartnerUpdateRequest) Validate() error {
	var errstrings []string

	if o.ReferenceID == "" {
		errstrings = append(errstrings, ValidationIsEmpty("reference id").Error())
	}

	if o.State == "" {
		errstrings = append(errstrings, ValidationIsEmpty("state").Error())
	}

	if err := o.Customer.Validate(); err != nil {
		errstrings = append(errstrings, o.Customer.Validate().Error())
	}

	if o.Offers == nil {
		errstrings = append(errstrings, ValidationIsEmpty("offers").Error())
	}

	if o.Offers != nil {
		for _, offer := range o.Offers {
			errstrings = append(errstrings, offer.Validate().Error())
		}
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

type FIO struct {
	FirstName  string `json:"firstName" bson:"first_name"`
	LastName   string `json:"lastName" bson:"last_name"`
	MiddleName string `json:"middleName" bson:"middle_name"`
}

func (f FIO) Validate() error {
	var errstrings []string

	if f.FirstName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("reference id").Error())
	}

	if f.LastName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("state").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
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

func (u UpdateMarketOrderRequest) Validate() error {
	var errstrings []string

	if u.ReferenceId == "" {
		errstrings = append(errstrings, ValidationIsEmpty("reference id").Error())
	}

	if u.State == "" {
		errstrings = append(errstrings, ValidationIsEmpty("state").Error())
	}

	if u.LoanLength == "" {
		errstrings = append(errstrings, ValidationIsEmpty("loan length").Error())
	}

	if u.Reason == "" {
		errstrings = append(errstrings, ValidationIsEmpty("reason").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
