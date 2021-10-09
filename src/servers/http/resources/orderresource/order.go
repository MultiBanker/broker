package orderresource

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
)

// @Tags ADMIN
// @Summary Получение заказа
// @Description Получение заказа
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "id of the order"
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.OrderRequest
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders/{id} [get]
func (o Order) order() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] empty ID")))
			return
		}

		res, err := o.orderMan.OrderByID(ctx, id)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
		}

		render.JSON(w, r, res)
		render.Status(r, http.StatusOK)

	}
}


// @Tags Techno-Order
// @Summary Создание нового заказа
// @Description Создание нового заказа
// @Accept  json
// @Produce  json
// @Param market body dto.OrderRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.IDResponse
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders [post]
func (o Order) neworder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.OrderRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		id, err := o.orderMan.NewOrder(ctx, &req)
		if err != nil {
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, &dto.BrokerResponse{
			ReferenceID: id,
		})
		render.Status(r, http.StatusOK)
	}
}

// @Tags Techno-Order
// @Summary Обновление заказа по решению клиента
// @Description Обновление заказа по решению клиента
// @Accept  json
// @Produce  json
// @Param market body dto.UpdateMarketOrderRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders/{reference_id}/partners [post]
func (o Order) updateClientOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.UpdateMarketOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	err := o.orderMan.UpdateMarketOrder(ctx, req)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
		render.Status(r, http.StatusOK)
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

}

// list godoc
// @Summary Получение заказов по reference_id
// @Description Получение заказов по reference_id
// @Tags Techno-Order
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param reference_id path string true "reference id of the order"
// @Success 200 {array} dto.OrderResponse
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders/{reference_id}/partners [get]
func (o Order) ordersByReference(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	referID := chi.URLParam(r, "reference_id")
	if referID == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("unauthorized")))
		return
	}

	marketCode, ok := ctx.Value("code").(string)
	if !ok {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("unauthorized")))
		return
	}

	orders, err := o.orderMan.PartnerOrder(ctx, marketCode, referID)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case errors.Is(err, nil):
		render.Status(r, http.StatusOK)
		render.JSON(w, r, orders)
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
}
