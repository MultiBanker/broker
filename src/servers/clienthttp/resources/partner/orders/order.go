package orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"
)

// @Tags Orders
// @Summary Обновление заказа по решению партнера
// @Description Обновление заказа по решению партнера
// @Accept  json
// @Produce  json
// @Param market body dto.OrderPartnerUpdateRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /partners/orders/ [post]
func (o Resource) updateorderpartner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.OrderPartnerUpdateRequest

		partnerCode, ok := ctx.Value("code").(string)
		if !ok {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] Unknown partner")))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		id, err := o.orderMan.UpdatePartnerOrder(ctx, models.PartnerOrder{
			ReferenceID: req.ReferenceID,
			State:       req.State,
			StateTitle:  req.StateTitle,
			PartnerCode: partnerCode,
			Customer: models.Customer{
				FirstName:  req.Customer.FirstName,
				MiddleName: req.Customer.MiddleName,
				LastName:   req.Customer.LastName,
			},
			Offers: req.Offers,
		})
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
