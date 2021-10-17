package offer

import (
	"encoding/json"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	_ "github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"

)

// @Tags Offers
// @Summary Получение офферов по заказу
// @Description Получение офферов по заказу
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param order body dto.OffersRequest true "body"
// @Success 200 {array} models.Offer
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /offers [post]
func (res Resource) offers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.OffersRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	offers, err := res.offerMan.OffersByGoods(ctx, req.Goods)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w,r, offers)
}
