package dto

import (
	"net/http"

	"github.com/go-chi/render"
)

type BrokerResponse struct {
	ReferenceID string `json:"reference_id"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type IDResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}
