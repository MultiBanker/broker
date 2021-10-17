package dto

import (
	"fmt"
	"strings"

	"github.com/MultiBanker/broker/src/models"
)

type OfferRequest struct {
	// Код партнера
	PartnerCode string `json:"partner_code" bson:"partner_code" example:"airba_pay"`
	// Наименование
	Name string `json:"name" bson:"name" example:"МФО Аирба"`
	// Вид оплаты товара
	PaymentTypeGroupCode string `json:"payment_type_group_code" bson:"payment_type_group_code" example:"online_broker"`
	// Минимальная сумма заказа
	MinOrderSum int `json:"min_order_sum" bson:"min_order_sum" example:"10000"`
	// Максимальная сумма заказа
	MaxOrderSum int `json:"max_order_sum" bson:"max_order_sum" example:"1500000"`
}

type OfferSpecs struct {
	Total  int64          `json:"total"`
	Offers []models.Offer `json:"offers"`
}

func (c OfferRequest) Validate() error {
	var errstrings []string

	if c.Name == "" {
		errstrings = append(errstrings, ValidationIsEmpty("name").Error())
	}
	if c.PartnerCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("partner code").Error())
	}
	if c.PaymentTypeGroupCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("payment type group code").Error())
	}
	if c.MinOrderSum == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("min order sum").Error())
	}

	if c.MaxOrderSum == 0 {
		errstrings = append(errstrings, ValidationIsEmpty("max order sum").Error())
	}

	if c.MinOrderSum > c.MaxOrderSum {
		errstrings = append(errstrings, fmt.Errorf("min order sum higher than max order sum").Error())
	}
	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
