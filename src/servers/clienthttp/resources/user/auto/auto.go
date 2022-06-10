package auto

import (
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	_ "github.com/MultiBanker/broker/src/models"
	"github.com/go-chi/render"
)

// @Summary      user auto
// @Description  user auto
// @Tags         User
// @Produce      json
// @Security     ApiTokenAuth
// @Success      200           {object} models.UserAuto
// @Failure      400           {object}  httperrors.Response
// @Failure      401           {object}  httperrors.Response
// @Failure      404           {object}  httperrors.Response
// @Failure      500           {object}  httperrors.Response
// @Router       /api/v1/users/auto [get]
func (res resource) getAuto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := ctx.Value("user_id").(string)

	userAuto, err := res.userAutoRepo.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, drivers.ErrDoesNotExist) {
			_ = render.Render(w, r, httperrors.ResourceNotFound(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, userAuto)
}
