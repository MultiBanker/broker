package dto

import (
	"net/http"

	"github.com/MultiBanker/broker/src/models"
	"github.com/go-chi/render"
)

type IDResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Markets struct {
	Total   int64           `json:"total"`
	Markets []models.Market `json:"markets"`
}

type Orders struct {
	Total  int64           `json:"total"`
	Orders []*models.Order `json:"orders"`
}

type Partners struct {
	Total    int64            `json:"total"`
	Partners []models.Partner `json:"partners"`
}

func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}
