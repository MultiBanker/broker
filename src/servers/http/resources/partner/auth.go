package partner

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/http/dto"

	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
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
// @Router /partners/login [post]
func (a Auth) auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.Login

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		partner, err := a.partnerMan.PartnerByUsername(ctx, req.Username, req.Password)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		access, refresh, err := a.authMan.Tokens(partner.ID, partner.Code, models.PARTNER)
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
// @Router /partners/logout [get]
func (a Auth) out() http.HandlerFunc {
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
