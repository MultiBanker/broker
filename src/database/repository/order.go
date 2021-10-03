package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) OrderRepo() OrderRepository {
	return r.Order
}

type OrderRepository interface {
	NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	OrderByID(ctx context.Context, id string) (dto.OrderRequest, error)
	Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, error)
	OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error)
}