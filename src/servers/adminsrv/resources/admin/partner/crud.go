package partner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/servers/adminsrv/dto"

	"github.com/MultiBanker/broker/src/models/selector"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
)

// @Tags Partner
// @Summary Создание нового партнера
// @Description Создание нового партнера
// @Accept  json
// @Produce  json
// @Param partner body models.Partner true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/partners [post]
func (a AdminResource) newpartner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.PartnerRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		partner := DTOToModelPartner(req)

		id, err := a.partnerMan.NewPartner(ctx, &partner)
		if err != nil {
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, &models.Response{ID: id, Status: "Created"})
		render.Status(r, http.StatusCreated)
	}
}

// @Tags Partner
// @Summary Обновление партнера
// @Description Обновление партнера
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param partner body dto.PartnerRequest true "body"
// @Param id path string true "id of the market"
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/partners/{id} [put]
func (a AdminResource) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] no id is passed")))
			return
		}

		var req dto.PartnerRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		partner := DTOToModelPartner(req)
		partner.ID = id

		id, err := a.partnerMan.UpdatePartner(ctx, &partner)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}
		render.JSON(w, r, &models.Response{ID: id, Status: "Updated"})
		render.Status(r, http.StatusCreated)
	}
}

// @Tags Partner
// @Summary Получение партнера
// @Description Получение партнера
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the market"
// @Success 200 {object} models.Partner
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/partners/{id} [get]
func (a AdminResource) partner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] no id is passed")))
			return
		}

		res, err := a.partnerMan.PartnerByID(ctx, id)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, res)
		render.Status(r, http.StatusOK)
	}
}

// @Tags Partner
// @Summary Получение партнеров
// @Description Получение партнеров
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success 200 {object} dto.Partners
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admins/partners [get]
func (a AdminResource) partners() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		res, total, err := a.partnerMan.Partners(ctx, &paging)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, dto.Partners{
			Total:    total,
			Partners: res,
		})
		render.Status(r, http.StatusOK)
	}
}
