package market

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

func DtoToModelMarket(req dto.MarketRequest) models.Market {
	return models.Market{
		CompanyName:    req.CompanyName,
		LogoURL:        req.LogoURL,
		Code:           req.Code,
		UpdateOrderURL: req.UpdateOrderURL,
		Username:       req.Username,
		Password:       &req.Password,
		Contact: models.ContactInfo{
			FirstName:  req.Contact.FirstName,
			LastName:   req.Contact.LastName,
			Email:      req.Contact.Email,
			Phone:      req.Contact.Phone,
			WebAddress: req.Contact.WebAddress,
			Location:   req.Contact.Location,
			BIN:        req.Contact.BIN,
		},
	}
}
