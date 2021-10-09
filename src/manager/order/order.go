package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/MultiBanker/broker/src/clients"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Order struct {
	orderColl        repository.Orderer
	partnerOrderColl repository.PartnerOrderer
	sequenceColl     repository.Sequencer
	partnerColl      repository.Partnerer
}

type Orderer interface {
	NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error)
	OrderByID(ctx context.Context, id string) (dto.OrderRequest, error)
	Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, int64, error)
	OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error)
	UpdateMarketOrder(ctx context.Context, req dto.UpdateMarketOrderRequest) error

	PartnerOrder(ctx context.Context, marketCode, referenceID string) ([]*dto.OrderResponse, error)
	UpdatePartnerOrder(ctx context.Context, req dto.OrderPartnerUpdateRequest) (string, error)
}

var _ Orderer = (*Order)(nil)

func NewOrder(repos repository.Repositories) Orderer {
	return Order{
		orderColl:        repos.OrderRepo(),
		partnerOrderColl: repos.PartnerOrderRepo(),
		partnerColl:      repos.PartnerRepo(),
		sequenceColl:     repos.SequenceRepo(),
	}
}

func (o Order) NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	idInt, err := o.sequenceColl.NextSequenceValue(ctx, models.OrderSequences)
	if err != nil {
		return "", err
	}
	order.ID = strconv.Itoa(idInt)
	id, err := o.orderColl.NewOrder(ctx, order)
	if err != nil {
		return "", err
	}

	wg := sync.WaitGroup{}
	for _, partnersCode := range order.PaymentPartners {
		wg.Add(1)
		go func(partnerCode string) {
			defer wg.Done()
			if err := o.BankOrder(ctx, id, partnerCode, *order); err != nil {
				log.Println(err)
			}
		}(partnersCode.Code)
	}
	wg.Wait()

	return id, nil
}

func (o Order) UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return o.orderColl.UpdateOrder(ctx, order)
}

func (o Order) OrderByID(ctx context.Context, id string) (dto.OrderRequest, error) {
	return o.orderColl.OrderByID(ctx, id)
}

func (o Order) Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, int64, error) {
	return o.orderColl.Orders(ctx, paging)
}

func (o Order) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error) {
	return o.orderColl.OrdersByReferenceID(ctx, referenceID)
}

func (o Order) PartnerOrder(ctx context.Context, partnerCode, referenceID string) ([]*dto.OrderResponse, error) {
	return o.partnerOrderColl.OrdersByReferenceID(ctx, partnerCode, referenceID)
}

func (o Order) UpdatePartnerOrder(ctx context.Context, req dto.OrderPartnerUpdateRequest) (string, error) {
	return o.partnerOrderColl.UpdateOrder(ctx, req)
}

func (o Order) UpdateMarketOrder(ctx context.Context, req dto.UpdateMarketOrderRequest) error {
	partner, err := o.partnerColl.PartnerByID(ctx, req.ProductCode)
	if err != nil {
		return err
	}

	err = o.orderColl.UpdateOrderState(ctx, req)
	if err != nil {
		return err
	}

	cli := clients.NewClient(partner.URL.Update, "")

	_, err = cli.RequestOrder(ctx, req, 3, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (o Order) BankOrder(ctx context.Context, id string, partnerCode string, order dto.OrderRequest) error {
	partner, err := o.partnerColl.PartnerByID(ctx, partnerCode)
	if err != nil {
		return fmt.Errorf("[ERROR] getting partner from db %v", err)
	}
	order.OrderID = id

	_, err = o.orderColl.NewOrder(ctx, &order)
	if err != nil {
		return fmt.Errorf("[ERROR] creating partner order %v", err)
	}

	bankCli := clients.NewClient(partner.URL.Create, "")
	b, err := bankCli.RequestOrder(ctx, order, 3, nil)
	if err != nil {
		return fmt.Errorf("[ERROR] requesting order to partner from url %v", err)
	}

	var orderResponse dto.OrderResponse

	if err = json.Unmarshal(b, &orderResponse); err != nil {
		return fmt.Errorf("[ERROR] unmarshalling order partner response from url %v", err)
	}

	log.Printf(
		"[INFO] OrderID - %s, Partner code - %s, status - %s, http code - %s, redirectURL - %s, uuid - %s, message - %s",
		id, partnerCode, orderResponse.Status, orderResponse.Code, orderResponse.RedirectURL, orderResponse.RequestUUID, orderResponse.RequestUUID,
	)

	orderResponse.PartnerCode = partnerCode
	orderResponse.State = models.INIT.Status()

	_, err = o.partnerOrderColl.NewOrder(ctx, orderResponse)
	return err
}
