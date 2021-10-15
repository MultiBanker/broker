package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/MultiBanker/broker/src/servers/clienthttp/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Tags Orders
// @Summary Создание нового заказа
// @Description Создание нового заказа
// @Accept  json
// @Produce  json
// @Param market body dto.MarketOrderRequest true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.IDResponse
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders [post]
func (o Resource) neworder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.MarketOrderRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}

		marketCode, ok := ctx.Value(middleware.CodeKey).(string)
		if !ok {
			_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] unauthorized")))
			return
		}

		id, err := o.orderMan.NewOrder(ctx, &models.Order{
			MarketCode:              marketCode,
			OrderState:              models.INIT.Status(),
			RedirectURL:             req.RedirectURL,
			Channel:                 req.Channel,
			StateCode:               models.INIT.Status(),
			ProductType:             req.ProductType,
			PaymentMethod:           req.PaymentMethod,
			IsDelivery:              req.IsDelivery,
			TotalCost:               req.Amount,
			LoanLength:              strconv.Itoa(req.LoanLength),
			VerificationId:          req.VerificationID,
			VerificationSMSCode:     req.VerificationSmsCode,
			VerificationSMSDatetime: req.VerificationSmsDateTime,
			Customer:                req.Customer,
			Address:                 req.Address,
			Goods:                   req.Goods,
			SystemCode:              req.SystemCode,
			PaymentPartners:         req.PaymentPartners,
		})
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

// @Tags Orders
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
// @Router /orders/markets [post]
//func (o Resource) marketOrderUpdate(w clienthttp.ResponseWriter, r *clienthttp.Request) {
//	ctx := r.Context()
//
//	var req dto.UpdateMarketOrderRequest
//
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		_ = render.Render(w, r, httperrors.BadRequest(err))
//		return
//	}
//
//	err := o.orderMan.UpdateMarketOrder(ctx, req)
//	switch {
//	case errors.Is(err, drivers.ErrDoesNotExist):
//		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
//		return
//	case errors.Is(err, nil):
//		render.Status(r, clienthttp.StatusOK)
//	default:
//		_ = render.Render(w, r, httperrors.Internal(err))
//		return
//	}
//
//}

// list godoc
// @Summary Получение заказов по reference_id
// @Description Получение заказов по reference_id
// @Tags Orders
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param reference_id path string true "reference id of the order"
// @Success 200 {array} models.Order
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /orders/{reference_id}/partners [get]
func (o Resource) ordersByReference(w http.ResponseWriter, r *http.Request) {
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
