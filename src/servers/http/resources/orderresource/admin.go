package orderresource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// list godoc
// @Summary Получение заказов
// @Description Получение заказов
// @Tags Order
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success 200 {object} dto.Orders
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders [get]
func (o Order) orders() http.HandlerFunc {
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

		res, total, err := o.orderMan.Orders(ctx, &paging)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
		}

		render.JSON(w, r, &dto.Orders{
			Total:  total,
			Orders: res,
		})
		render.Status(r, http.StatusOK)

	}
}

// @Tags Order
// @Summary Обновление заказа
// @Description Обновление заказа
// @Accept  json
// @Produce  json
// @Param market body dto.OrderRequest true "body"
// @Security ApiKeyAuth
// @Param id path string true "id of the order"
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders/{id} [put]
func (o Order) updateorder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] ID Order is empty")))
			return
		}

		var req models.Order

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		id, err := o.orderMan.UpdateOrder(ctx, &req)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
		}

		render.JSON(w, r, &models.Response{
			ID:     id,
			Status: "updated",
		})
		render.Status(r, http.StatusOK)
	}
}