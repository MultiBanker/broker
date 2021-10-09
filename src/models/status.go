package models

type OrderStatus int

const (
	INIT OrderStatus = iota + 1
	SIGNED
	APPROVED
	PREAPPROVED
	RESERVED
	REJECTED
	FINANCED
	CONTRACTCANCELLED
	DELIVERED
	PARTNERCANCELLED
	CUSTOMERAPPROVED
	CANCELLED
)

var (
	status = map[OrderStatus]string{
		INIT:              "INIT",
		SIGNED:            "SIGNED",
		APPROVED:          "APPROVED",
		PREAPPROVED:       "PREAPPROVED",
		RESERVED:          "RESERVED",
		REJECTED:          "REJECTED",
		FINANCED:          "FINANCED",
		CONTRACTCANCELLED: "CONTRACTCANCELLED",
		DELIVERED:         "DELIVERED",
		PARTNERCANCELLED:  "PARTNERCANCELLED",
		CUSTOMERAPPROVED:  "CUSTOMERAPPROVED",
		CANCELLED:         "CANCELLED",
	}
	title = map[OrderStatus]string{
		INIT:              "заявка создана",
		SIGNED:            "подписана",
		APPROVED:          "заявка одобрена банком",
		PREAPPROVED:       "PREAPPROVED",
		RESERVED:          "RESERVED",
		REJECTED:          "REJECTED",
		FINANCED:          "FINANCED",
		CONTRACTCANCELLED: "CONTRACTCANCELLED",
		DELIVERED:         "DELIVERED",
		PARTNERCANCELLED:  "PARTNERCANCELLED",
		CUSTOMERAPPROVED:  "CUSTOMERAPPROVED",
		CANCELLED:         "CANCELLED",
	}
)

func (o OrderStatus) Status() string {
	return status[o]
}

func (o OrderStatus) Title() string {
	return title[o]
}
