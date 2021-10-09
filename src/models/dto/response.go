package dto

import (
	"net/http"

	"github.com/go-chi/render"
)

type BrokerResponse struct {
	ReferenceID string `json:"reference_id"`
}

func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}
