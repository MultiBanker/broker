package order

import (
	"context"
	"sync"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/banker"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Order struct {
	orderColl    repository.OrderRepository
	sequenceColl repository.SequencesRepository
	banker       []banker.Banker
}

type Orderer interface {
	NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	OrderByID(ctx context.Context, id string) (dto.OrderRequest, error)
	Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, error)
	OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error)
}

func NewOrder(repos repository.Repositories, banker ...banker.Banker) Orderer {
	return Order{
		orderColl:    repos.OrderRepo(),
		sequenceColl: repos.SequenceRepo(),
		banker:       banker,
	}
}

func (o Order) NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	_, err := o.orderColl.NewOrder(ctx, order)
	if err != nil {
		return "", err
	}
	wg := sync.WaitGroup{}
	for _, bank := range o.banker {
		wg.Add(1)
		go func(bank banker.Banker) {
			defer wg.Done()
			bank.CreateOrder(ctx, *order)
		}(bank)
	}
	wg.Wait()

	return order.ReferenceID, nil
}

func (o Order) UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return o.orderColl.UpdateOrder(ctx, order)
}

func (o Order) OrderByID(ctx context.Context, id string) (dto.OrderRequest, error) {
	return o.orderColl.OrderByID(ctx, id)
}

func (o Order) Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, error) {
	return o.orderColl.Orders(ctx, paging)
}

func (o Order) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error) {
	return o.orderColl.OrdersByReferenceID(ctx, referenceID)
}