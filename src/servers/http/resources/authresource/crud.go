package authresource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models/selector"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
)

func (a Auth) newpartner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req models.Partner

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		id, err := a.partnerMan.NewPartner(ctx, &req)
		if err != nil {
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, &models.Response{ID: id, Status: "Created"})
		render.Status(r, http.StatusCreated)
	}
}

func (a Auth) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")
		if id == "" {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] no id is passed")))
			return
		}

		var req models.Partner

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		id, err := a.partnerMan.UpdatePartner(ctx, &req)
		if err != nil {
			_ = render.Render(w, r, httperrors.Internal(err))
			return
		}

		render.JSON(w, r, &models.Response{ID: id, Status: "Updated"})
		render.Status(r, http.StatusCreated)
	}
}

func (a Auth) partner() http.HandlerFunc {
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

func (a Auth) partners() http.HandlerFunc {
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

		res, err := a.partnerMan.Partners(ctx, &paging)
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

