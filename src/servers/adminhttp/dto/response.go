package dto

import (
	"net/http"

	"github.com/MultiBanker/broker/src/models"
	"github.com/go-chi/render"
)

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

type LoanPrograms struct {
	Total        int64                `json:"total"`
	LoanPrograms []models.LoanProgram `json:"loan_programs"`
}

func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}
