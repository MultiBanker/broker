package partner

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

func DTOToModelPartner(req dto.PartnerRequest) models.Partner {
	return models.Partner{
		Code:        req.Code,
		CompanyName: req.CompanyName,
		Commission:  req.Commission,
		LogoURL:     req.LogoURL,
		URL: &models.URL{
			Create: req.URL.Create,
			Update: req.URL.Update,
		},
		Contact: models.ContactInfo{
			FirstName:  req.Contact.FirstName,
			LastName:   req.Contact.LastName,
			Email:      req.Contact.Email,
			Phone:      req.Contact.Phone,
			WebAddress: req.Contact.WebAddress,
			Location:   req.Contact.Location,
			BIN:        req.Contact.BIN,
		},
		Enabled:  req.Enabled,
		Username: req.Username,
		Password: req.Password,
	}
}
