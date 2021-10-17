package dto

import (
	"fmt"
	"strings"

	"github.com/MultiBanker/broker/src/models"
)

type OffersRequest struct {
	Goods []*models.Goods `json:"goods"`
}

func (o OffersRequest) Validate() error {
	var errstrings []string
	if o.Goods == nil {
		errstrings = append(errstrings, ValidationIsEmpty("goods").Error())
	}

	if o.Goods != nil {
		for _, good := range o.Goods {
			errstrings = append(errstrings, good.Validate().Error())
		}
	}
	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
