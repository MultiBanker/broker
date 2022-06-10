package application

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	_ "github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	_ "github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        auto       body      dto.UserApplication  true  "add application"
// @Success      200  {object}  models.Response
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/ [post]
func (res resource) createApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.Context().Value("user_id").(string)

	var req dto.UserApplication

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.user.Create(ctx, models.UserApplication{
		UserID:    userID,
		ChosenSKU: req.ChosenSKU,
	})
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, models.Response{
		ID:     id,
		Status: "created",
	})

}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        id  path      string  true  "application id"
// @Success      200           {object} models.UserApplication
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/{id} [get]
func (res resource) getApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	appl, err := res.user.Get(ctx, id)
	if err != nil {
		if errors.Is(err, drivers.ErrDoesNotExist) {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, appl)
}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success      200           {object} dto.UserApplications
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/list [get]
func (res resource) listApplications(w http.ResponseWriter, r *http.Request) {

}

// @Summary      user application
// @Description  user application
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param        id  path      string  true  "application id"
// @Success      204
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/application/{id} [delete]
func (res resource) deleteApplication(w http.ResponseWriter, r *http.Request) {

}
