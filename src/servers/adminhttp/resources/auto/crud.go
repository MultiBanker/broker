package auto

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Summary      add car product
// @Description  add car product
// @Tags         Auto
// @Produce      json
// @Security     ApiTokenAuth
// @Param        auto       body      dto.Auto  true  "add car"
// @Success      200           {object}  models.Response
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/auto [post]
func (res Resource) addCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.Auto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	sku, err := res.autoMan.Create(ctx, DtoToAuto(req))
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	dto.RespondJSON(w, r, http.StatusOK, models.Response{
		ID:     sku,
		Status: "created",
	})

}

// @Summary      update car product
// @Description  update car product
// @Tags         Auto
// @Produce      json
// @Security     ApiTokenAuth
// @Param        sku  path      string  true  "sku"
// @Param        auto       body      dto.Auto  true  "add car"
// @Success      200           {object}  models.Response
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/auto/{sku} [put]
func (res Resource) updateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.Auto

	sku := chi.URLParam(r, "sku")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	autoModel := DtoToAuto(req)
	autoModel.SKU = sku

	sku, err := res.autoMan.Update(ctx, autoModel)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	dto.RespondJSON(w, r, http.StatusOK, models.Response{
		ID:     sku,
		Status: "updated",
	})
}

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

// @Summary      delete car product
// @Description  delete car product
// @Tags         Auto
// @Produce      json
// @Security     ApiTokenAuth
// @Param        sku  path      string  true  "sku"
// @Success      204
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/auto/{sku} [delete]
func (res Resource) deleteCar(w http.ResponseWriter, r *http.Request) {
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
// @Success      200           {object}  dto.Auto
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
