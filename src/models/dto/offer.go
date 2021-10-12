package dto

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/pkg/errors"
)

type OfferRequest struct {
	PartnerCode          string `json:"partner_code" bson:"partner_code"`
	Name                 string `json:"name" bson:"name"`
	PaymentTypeGroupCode string `json:"payment_type_group_code" bson:"payment_type_group_code"`
	MinOrderSum          int    `json:"min_order_sum" bson:"min_order_sum"`
	MaxOrderSum          int    `json:"max_order_sum" bson:"max_order_sum"`
}

type OfferSpecs struct {
	Total  int64
	Offers []models.Offer
}

func (c OfferRequest) Validate() error {
	if c.Name == "" {
		return errors.Wrap(IsEmpty, "name")
	}
	if c.PartnerCode == "" {
		return errors.Wrap(IsEmpty, "partner code")
	}
	if c.PaymentTypeGroupCode == "" {
		return errors.Wrap(IsEmpty, "payment type group code")
	}
	if c.MinOrderSum == 0 {
		return errors.Wrap(IsEmpty, "min order sum")
	}

	if c.MaxOrderSum == 0 {
		return errors.Wrap(IsEmpty, "max order sum")
	}

	return nil
}
