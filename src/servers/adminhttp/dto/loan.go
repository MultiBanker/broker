package dto

import (
	"fmt"

	"github.com/MultiBanker/broker/src/models"
)

type LoanProgramRequest struct {
	Code        string  `json:"code" bson:"code"`
	IsEnabled   bool    `json:"is_enabled" bson:"is_enabled"`
	MaxAmount   int     `json:"max_amount" bson:"max_amount"`
	MinAmount   int     `json:"min_amount" bson:"min_amount"`
	Name        string  `json:"name" bson:"name"`
	Note        string  `json:"note" bson:"note"`
	PartnerCode string  `json:"partner_code" bson:"partner_code"`
	Rate        float64 `json:"rate" bson:"rate"`
	Term        int     `json:"term" bson:"term"`
	Type        string  `json:"type" bson:"type"`
}

func (l LoanProgramRequest) Validate() error {
	if l.Type != models.INSTALLMENT.String() || l.Type != models.LOAN.String() {
		return fmt.Errorf("[ERROR] invalid type")
	}

	if l.Name == "" {
		return fmt.Errorf("[ERROR] empty name")
	}

	if l.Code == "" {
		return fmt.Errorf("[ERROR] empty code")
	}

	if l.Note == "" {
		return fmt.Errorf("[ERROR] empty note")
	}

	if l.PartnerCode == "" {
		return fmt.Errorf("[ERROR] invalid partner code")
	}

	if l.Term == 0 {
		return fmt.Errorf("[ERROR] empty loan term")
	}

	if l.MinAmount > l.MaxAmount || l.MinAmount == 0 || l.MaxAmount == 0 {
		return fmt.Errorf("[ERROR] wrong min and max amount")
	}

	if l.Rate == 0 {
		return fmt.Errorf("[ERROR] wrong rate")
	}

	return nil
}
