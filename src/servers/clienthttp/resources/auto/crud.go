package auto

import (
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	_ "github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Summary      get car product
// @Description  get car product
// @Tags         Auto
// @Produce      json
// @Security     ApiTokenAuth
// @Param        sku  path      string  true  "sku"
// @Success      200           {object}  models.Auto
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/auto/{sku} [get]
func (res Resource) getCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sku := chi.URLParam(r, "sku")

	auto, err := res.autoMan.Get(ctx, sku)
	if err != nil {
		if errors.Is(err, drivers.ErrDoesNotExist) {
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	dto.RespondJSON(w, r, http.StatusOK, auto)
}

// @Summary      list car product
// @Description  list car product
// @Tags         Auto
// @Produce      json
// @Security     ApiTokenAuth
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success      200           {object}  dto.ListAuto
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/auto [get]
func (res Resource) listCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sq, err := selector.NewSearchQuery(r.URL.Query())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	auto, count, err := res.autoMan.List(ctx, *sq)
	if err != nil {
		if errors.Is(err, drivers.ErrDoesNotExist) {
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	dto.RespondJSON(w, r, http.StatusOK, dto.ListAuto{
		Autos: auto,
		Count: count,
	})
}
