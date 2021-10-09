package orderresource

import (
	"fmt"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
)

// @Tags Order
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


