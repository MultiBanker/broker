package dto

import (
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
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

func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}

func RespondForbidden(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusForbidden)
	_ = render.Render(w, r, httperrors.AccessDenied(err))
}
