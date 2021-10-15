package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) OrderRepo() Orderer {
	return r.Order
}

func (r Repository) PartnerOrderRepo() PartnerOrderer {
	return r.PartnerOrder
}

type Orderer interface {
	NewOrder(ctx context.Context, order *models.Order) (string, error)
	UpdateOrder(ctx context.Context, order *models.Order) (string, error)
	OrderByID(ctx context.Context, id string) (models.Order, error)
	Orders(ctx context.Context, paging *selector.Paging) ([]*models.Order, int64, error)
	OrdersByReferenceID(ctx context.Context, referenceID string) ([]*models.Order, error)
	//UpdateOrderState(ctx context.Context, request dto.UpdateMarketOrderRequest) error
}

type PartnerOrderer interface {
	NewOrder(ctx context.Context, order models.PartnerOrder) (string, error)
	UpdateOrder(ctx context.Context, order models.PartnerOrder) (string, error)
	OrdersByReferenceID(ctx context.Context, marketCode, referenceID string) ([]*models.PartnerOrder, error)
	OrderPartner(ctx context.Context, referenceID, partnerCode string) (models.PartnerOrder, error)
	UpdateInitStatusByTimeOut(ctx context.Context) error
}
