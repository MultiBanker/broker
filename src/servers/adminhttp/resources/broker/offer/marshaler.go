package offer

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

func DtoToModelOffer(req dto.OfferRequest) models.Offer {
	return models.Offer{
		Name:                 req.Name,
		PartnerCode:          req.PartnerCode,
		PaymentTypeGroupCode: req.PaymentTypeGroupCode,
		MinOrderSum:          req.MinOrderSum,
		MaxOrderSum:          req.MaxOrderSum,
	}
}
