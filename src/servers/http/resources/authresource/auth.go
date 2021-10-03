package authresource

import (
	"encoding/json"
	"github.com/MultiBanker/broker/pkg/httperrors"
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models/dto"
)

const (
	tokenTTL = 60 * time.Minute
)

func (a Auth) auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.Login

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		partner, err := a.partnerMan.PartnerByUsername(ctx, req.Username)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
		}

		access, _, err := a.authMan.Tokens(partner.ID)
		if err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		render.JSON(w, r, &dto.TokenResponse{AccessToken: access})
		render.Status(r, http.StatusOK)
	}
}

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