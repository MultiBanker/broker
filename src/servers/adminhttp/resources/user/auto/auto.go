package auto

import (
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	"github.com/go-chi/render"
)

// @Summary      user auto
// @Description  user auto
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success      200           {object} dto.ListUserAuto
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/auto [get]
func (res resource) getAuto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sq, err := selector.NewSearchQuery(r.URL.Query())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	userAuto, count, err := res.userAutoRepo.List(ctx, *sq)
	if err != nil {
		if errors.Is(err, drivers.ErrDoesNotExist) {
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.ListUserAuto{
		Autos: userAuto,
		Count: count,
	})
}
