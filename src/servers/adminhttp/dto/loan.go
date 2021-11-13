package dto

import (
	"fmt"

	"github.com/MultiBanker/broker/src/models"
)

type LoanProgramRequest struct {
	// Код продукта
	Code        string  `json:"code" example:"001"`
	// Доступность
	IsEnabled   bool    `json:"is_enabled" example:"true"`
	// Максимальная сумма корзины
	MaxAmount   int     `json:"max_amount" example:"10000000"`
	// Минимальная сумма корзины
	MinAmount   int     `json:"min_amount" example:"5000"`
	// Наименование продукта
	Name        string  `json:"name" example:"Installment 1"`
	// Примечание
	Note        string  `json:"note" example:"Note"`
	// Код банка партнера
	PartnerCode string  `json:"partner_code" example:"mfo_ff"`
	// Ставка по кредиту
	Rate        float64 `json:"rate" example:"0.3"`
	// Срок
	Term        int     `json:"term" example:"12"`
	// Тип
	Type        string  `json:"type" example:"installment" enums:"loan,installment"`
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
