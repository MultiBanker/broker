package offer

import (
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/http/dto"
	"github.com/go-chi/render"
)

const maxOrderHistoryLimit = 100

// @Tags Offers
// @Summary Получение офферов
// @Description Получение офферов
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success 200 {object} dto.OfferSpecs
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /offers [get]
func (res Resource) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// строим пагинацию
	paging := selector.Paging{
		SortKey: "created_at",
		SortVal: -1,
	}
	skipStr := r.URL.Query().Get("skip")
	skip, err := strconv.ParseInt(skipStr, 10, 64)
	if err != nil {
		skip = 0
	}

	if skip > 0 {
		paging.Skip = skip
	}
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	if limit > 0 && limit < maxOrderHistoryLimit {
		paging.Limit = limit
	}

	offers, total, err := res.offerMan.Offers(ctx, paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.OfferSpecs{
		Total:  total,
		Offers: offers,
	})
}
