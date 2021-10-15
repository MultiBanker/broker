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
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
)

type Order struct {
	orderColl        repository.Orderer
	partnerOrderColl repository.PartnerOrderer
	sequenceColl     repository.Sequencer
	marketColl       repository.Marketer
	partnerColl      repository.Partnerer
}

type Orderer interface {
	NewOrder(ctx context.Context, order *models.Order) (string, error)
	UpdateOrder(ctx context.Context, order *models.Order) (string, error)
	OrderByID(ctx context.Context, id string) (models.Order, error)
	Orders(ctx context.Context, paging *selector.Paging) ([]*models.Order, int64, error)
	OrdersByReferenceID(ctx context.Context, referenceID string) ([]*models.Order, error)
	//UpdateMarketOrder(ctx context.Context, req models.Order) error

	PartnerOrder(ctx context.Context, marketCode, referenceID string) ([]*models.PartnerOrder, error)
	UpdatePartnerOrder(ctx context.Context, req models.PartnerOrder) (string, error)
}

var _ Orderer = (*Order)(nil)

func NewOrder(repos repository.Repositories) Orderer {
	return Order{
		orderColl:        repos.OrderRepo(),
		partnerOrderColl: repos.PartnerOrderRepo(),
		partnerColl:      repos.PartnerRepo(),
		sequenceColl:     repos.SequenceRepo(),
		marketColl:       repos.MarketRepo(),
	}
}

func (o Order) NewOrder(ctx context.Context, order *models.Order) (string, error) {
	idInt, err := o.sequenceColl.NextSequenceValue(ctx, models.OrderSequences)
	if err != nil {
		return "", err
	}
	order.ID = strconv.Itoa(idInt)
	id, err := o.orderColl.NewOrder(ctx, order)
	if err != nil {
		return "", err
	}

	order.ReferenceID = id

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

func (o Order) UpdateOrder(ctx context.Context, order *models.Order) (string, error) {
	return o.orderColl.UpdateOrder(ctx, order)
}

func (o Order) OrderByID(ctx context.Context, id string) (models.Order, error) {
	return o.orderColl.OrderByID(ctx, id)
}

func (o Order) Orders(ctx context.Context, paging *selector.Paging) ([]*models.Order, int64, error) {
	return o.orderColl.Orders(ctx, paging)
}

func (o Order) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*models.Order, error) {
	return o.orderColl.OrdersByReferenceID(ctx, referenceID)
}

func (o Order) PartnerOrder(ctx context.Context, marketCode, referenceID string) ([]*models.PartnerOrder, error) {
	return o.partnerOrderColl.OrdersByReferenceID(ctx, marketCode, referenceID)
}

func (o Order) UpdatePartnerOrder(ctx context.Context, req models.PartnerOrder) (string, error) {
	_, err := o.partnerOrderColl.UpdateOrder(ctx, req)
	if err != nil {
		return "", err
	}
	order, err := o.partnerOrderColl.OrderPartner(ctx, req.ReferenceID, req.PartnerCode)
	if err != nil {
		return "", err
	}
	m, err := o.marketColl.MarketByCode(ctx, order.MarketCode)
	if err != nil {
		return "", err
	}
	cli := clients.NewClient(m.UpdateOrderURL, "")
	_, err = cli.RequestOrder(ctx, req, 3, nil)
	return "", err
}

//func (o Order) UpdateMarketOrder(ctx context.Context, req models.Order) error {
//	partner, err := o.partnerColl.PartnerByCode(ctx, req.ProductCode)
//	if err != nil {
//		return err
//	}
//
//	//err = o.orderColl.UpdateOrderState(ctx, req)
//	//if err != nil {
//	//	return err
//	//}
//
//	cli := clients.NewClient(partner.URL.Update, "")
//
//	_, err = cli.RequestOrder(ctx, req, 3, nil)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}

func (o Order) BankOrder(ctx context.Context, id string, partnerCode string, order models.Order) error {
	partner, err := o.partnerColl.PartnerByCode(ctx, partnerCode)
	if err != nil {
		return fmt.Errorf("[ERROR] getting partner from db %v", err)
	}

	bankCli := clients.NewClient(partner.URL.Create, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzeXN0ZW1fY29kZSI6IlREX0JST0tFUiJ9.Fa4wL9KID3A_-8fYmvhZKXi68K5GRMlLsYK0y6PASI4")
	b, err := bankCli.RequestOrder(ctx, OrderToDTO(order), 3, nil)
	if err != nil {
		return fmt.Errorf("[ERROR] requesting order to partner from url %v", err)
	}

	var orderResponse dto.PartnerOrderResponse

	if err = json.Unmarshal(b, &orderResponse); err != nil {
		return fmt.Errorf("[ERROR] unmarshalling order partner response from url %v", err)
	}

	log.Printf(
		"[INFO] OrderID - %s, Partner code - %s, status - %s, clienthttp code - %d, redirectURL - %s, uuid - %s, message - %s",
		id, partnerCode, orderResponse.Status, orderResponse.Code, orderResponse.RedirectUrl, orderResponse.RequestUuid, orderResponse.Message,
	)

	_, err = o.partnerOrderColl.NewOrder(ctx, models.PartnerOrder{
		ReferenceID: order.ReferenceID,
		PartnerCode: partnerCode,
		MarketCode:  order.MarketCode,
		Status:      orderResponse.Status,
		Code:        orderResponse.Code,
		RedirectURL: orderResponse.RedirectUrl,
		RequestUUID: orderResponse.RequestUuid,
		Message:     orderResponse.Message,
		State:       models.INIT.Status(),
		StateTitle:  models.INIT.Title(),
		Offers:      []models.Offers{},
	})
	return err
}

func OrderToDTO(order models.Order) dto.PartnerOrderRequest {
	return dto.PartnerOrderRequest{
		ReferenceId:             order.ReferenceID,
		IsDelivery:              order.IsDelivery,
		Channel:                 order.Channel,
		PaymentMethod:           order.PaymentMethod,
		ProductType:             order.ProductType,
		RedirectUrl:             order.RedirectURL,
		OrderState:              order.OrderState,
		SalesPlace:              order.SalesPlace,
		VerificationSmsCode:     order.VerificationSMSCode,
		VerificationSmsDateTime: order.VerificationSMSDatetime,
		LoanLength:              order.LoanLength,
		Customer:                order.Customer,
		Address:                 order.Address,
		Goods:                   order.Goods,
		TotalCost:               order.TotalCost,
	}
}
