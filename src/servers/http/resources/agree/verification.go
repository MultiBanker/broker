package agree

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/manager/signature"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Tags Signature
// @Summary Валидация токена
// @Description Валидация токена
// @Accept  json
// @Produce  json
// @Param token body dto.CheckVerificationReq true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the signature"
// @Success 200
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /signatures/{id} [put]
func (res Resource) token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CheckVerificationReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	req.SignID = chi.URLParam(r, "id")
	if req.SignID == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] empty sign id")))
		return
	}

	_, err := res.signatureMan.CheckVerification(ctx, req.SignID, req.Token)
	switch err {
	case drivers.ErrDoesNotExist:
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case nil:
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
}

// @Tags Signature
// @Summary Переотправка токена
// @Description Переотправка токена
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the signature"
// @Success 200
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /signatures/{id}/resend [patch]
func (res Resource) resend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	signID := chi.URLParam(r, "id")
	if signID == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("invalid sign ID")))
		return
	}

	id, err := res.signatureMan.Update(ctx, signID)
	switch err {
	case signature.ErrTryLimitReached:
		_ = render.Render(w, r, httperrors.AccessDenied(err))
		return
	case nil:

	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.IDResponse{
		ID:     id,
		Status: "updated",
	})
}

