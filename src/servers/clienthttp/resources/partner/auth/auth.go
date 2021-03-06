package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"
)

const (
	tokenTTL = 60 * time.Minute
)

// @Tags Partner
// @Summary Авторизация партнера
// @Description Авторизация партнера
// @Accept  json
// @Produce  json
// @Param auth body dto.Login true "body"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/partners/login [post]
func (res Resource) auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.Login

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		partnerResult, err := res.partnerMan.PartnerByUsername(ctx, req.Username, req.Password)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:

		case partner.ErrAuthorization:
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		access, refresh, err := res.authMan.Tokens(partnerResult.ID, partnerResult.Code, models.PARTNER)
		if err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		render.JSON(w, r, &dto.TokenResponse{AccessToken: access, ResponseToken: refresh})
		render.Status(r, http.StatusOK)
	}
}

// @Tags Partner
// @Summary выход авторизации партнера
// @Description выход авторизации партнера
// @Accept  json
// @Produce  json
// @Router /api/v1/partners/logout [get]
func (res Resource) out() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "Authorization",
			Value:    "",
			HttpOnly: false,
			Expires:  time.Now().In(time.UTC).Add(-tokenTTL),
		})

		render.Status(r, http.StatusOK)
	}
}
