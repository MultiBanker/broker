package dto

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/hashicorp/go-multierror"
)

type Auto struct {
	Title models.LangOptions `json:"title" bson:"title"`
	Brand models.Brand       `json:"brand" bson:"brand"`
	Color models.Color       `json:"color" bson:"color"`
	Media models.Medias      `json:"media" bson:"media"`
	About models.LangOptions `json:"about" bson:"about"`
	Price models.Price       `json:"price" bson:"price"`
}

func (a Auto) Validate() error {
	var result *multierror.Error

	if err := a.Title.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := a.Brand.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := a.Media.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := a.Color.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := a.Price.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := a.About.Validate(); err != nil {
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

type ListAuto struct {
	Autos []models.Auto `json:"autos"`
	Count int64         `json:"count"`
}

type ConnectAuto struct {
	SKU string
	VIN string
}

type ListUserAuto struct {
	Autos []models.UserAuto `json:"autos"`
	Count int64             `json:"count"`
}
