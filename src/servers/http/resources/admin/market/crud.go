package market

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/http/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const maxOrderHistoryLimit = 100

// @Tags Market
// @Summary Создание нового маркета
// @Description Создание нового маркета
// @Accept  json
// @Produce  json
// @Param market body dto.MarketRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/markets [post]
func (res AdminResource) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.MarketRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.market.CreateMarket(ctx, DtoToModelMarket(req))
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.JSON(w, r, &models.Response{
		ID:     id,
		Status: "created",
	})
	render.Status(r, http.StatusOK)
}

// @Tags Market
// @Summary Получение маркета
// @Description Получение маркета
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "id of the partner"
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Market
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/markets/{id} [get]
func (res AdminResource) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] Empty ID")))
		return
	}

	mm, err := res.market.MarketByID(ctx, id)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.JSON(w, r, mm)
	render.Status(r, http.StatusOK)
}

// list godoc
// @Summary Получение маркетов
// @Description Получение маркетов
// @Tags Market
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param page query int false "pagination skip"
// @Success 200 {object} dto.Markets
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/markets [get]
func (res AdminResource) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := selector.Paging{
		SortKey: "created_at",
		SortVal: -1,
	}
	q := r.URL.Query()
	skipStr := q.Get("page")
	skip, err := strconv.ParseInt(skipStr, 10, 64)
	if err != nil {
		skip = 0
	}

	if skip > 0 {
		paging.Skip = skip
	}
	limitStr := q.Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	if limit > 0 && limit < maxOrderHistoryLimit {
		paging.Limit = limit
	}

	mm, total, err := res.market.Markets(ctx, paging)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.JSON(w, r, &dto.Markets{
		Total:   total,
		Markets: mm,
	})
	render.Status(r, http.StatusOK)
}

// list godoc
// @Summary Обновление определенного маркета
// @Description Обновление определенного маркета
// @Tags Market
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param market body dto.MarketRequest true "body"
// @Param id path string true "id of the partner"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/markets/{id} [put]
func (res AdminResource) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.MarketRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] Empty ID")))
		return
	}

	market := DtoToModelMarket(req)
	market.ID = id

	id, err := res.market.UpdateMarket(ctx, market)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.JSON(w, r, &models.Response{
		ID:     id,
		Status: "updated",
	})
	render.Status(r, http.StatusOK)
}
