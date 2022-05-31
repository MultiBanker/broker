package auto

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

func DtoToAuto(req dto.Auto) models.Auto {
	return models.Auto{
		Title: req.Title,
		Brand: req.Brand,
		Color: req.Color,
		Media: req.Media,
		About: req.About,
		Price: req.Price,
	}
}
