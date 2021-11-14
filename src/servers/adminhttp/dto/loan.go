package dto

import (
	"fmt"
	"strings"

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
	var errstrings []string

	if _, err := models.ValidateProductType(l.Type); err != nil {
		errstrings = append(errstrings, fmt.Errorf("invalid type").Error())
	}
	if l.Name == "" {
		errstrings = append(errstrings, ValidationIsEmpty("name").Error())
	}

	if l.Code == "" {
		errstrings = append(errstrings, ValidationIsEmpty("code").Error())
	}

	if l.Note == "" {
		errstrings = append(errstrings, ValidationIsEmpty("note").Error())
	}

	if l.PartnerCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("partner code").Error())
	}

	if l.Term == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("term").Error())
	}

	if l.MinAmount > l.MaxAmount || l.MinAmount == 0 || l.MaxAmount == 0 {
		errstrings = append(errstrings, fmt.Errorf("[ERROR] wrong min and max amount").Error())
	}

	if l.Rate == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("rate").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, " and "))
	}
	return nil
}
