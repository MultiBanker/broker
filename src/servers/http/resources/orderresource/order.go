package orderresource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models/selector"

	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
)

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

		render.JSON(w, r, &models.Response{
			ID:     id,
			Status: "created",
		})
		render.Status(r, http.StatusOK)
	}
}

func (o Order) updateorder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] ID Order is empty")))
			return
		}

		var req dto.OrderRequest

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

		res, err := o.orderMan.Orders(ctx, &paging)
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

func (o Order) referenceOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] empty reference ID")))
			return
		}

		orders, err := o.orderMan.OrdersByReferenceID(ctx, id)
		switch err {
		case drivers.ErrDoesNotExist:
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		case nil:
		default:
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, orders)
		render.Status(r, http.StatusOK)
	}
}


