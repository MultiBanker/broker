package dto

import "github.com/MultiBanker/broker/src/models"

type ListAuto struct {
	Autos []models.Auto `json:"autos"`
	Count int64         `json:"count"`
}
