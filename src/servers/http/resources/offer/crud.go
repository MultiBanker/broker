package offer

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

const maxOrderHistoryLimit = 100

// @Tags Offers
// @Summary Создание нового оффера
// @Description Создание нового оффера
// @Accept  json
// @Produce  json
// @Param partner body dto.OfferRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.IDResponse
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /offers [post]
func (res Resource) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.OfferRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.offerMan.CreateOffer(ctx, models.Offer{
		Name:                 req.Name,
		PartnerCode:          req.PartnerCode,
		PaymentTypeGroupCode: req.PaymentTypeGroupCode,
		MinOrderSum:          req.MinOrderSum,
		MaxOrderSum:          req.MaxOrderSum,
	})
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.IDResponse{
		ID:     id,
		Status: "created",
	})
}

// @Tags Offers
// @Summary Обновление оффера
// @Description Обновление оффера
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param partner body dto.OfferRequest true "body"
// @Param id path string true "id of the market"
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Offer
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /offers/{id} [put]
func (res Resource) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		_ = render.Render(w, r, httperrors.BadRequest(errors.Wrap(dto.IsEmpty, "id")))
		return
	}
	var req dto.OfferRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	offer, err := res.offerMan.UpdateOffer(ctx, models.Offer{
		Name:                 req.Name,
		PartnerCode:          req.PartnerCode,
		PaymentTypeGroupCode: req.PaymentTypeGroupCode,
		MinOrderSum:          req.MinOrderSum,
		MaxOrderSum:          req.MaxOrderSum,
	})
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, offer)
}

// @Tags Offers
// @Summary Получение оффера
// @Description Получение оффера
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param code path string true "id of the market"
// @Success 200 {object} models.Partner
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /offers/{id} [get]
func (res Resource) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := chi.URLParam(r, "code")
	if code == "" {
		_ = render.Render(w, r, httperrors.BadRequest(errors.Wrap(dto.IsEmpty, "code")))
		return
	}

	offer, err := res.offerMan.OfferByCode(ctx, code)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, offer)
}

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
