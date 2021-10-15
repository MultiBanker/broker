package models

import "fmt"

var ErrUnknownProductType = fmt.Errorf("[ERROR] unknown product type")

type ProductType int

const (
	INSTALLMENT ProductType = iota + 1
	LOAN
)

var (
	productType = map[ProductType]string{
		INSTALLMENT: "installment",
		LOAN:        "loan",
	}

	validateProductType = map[string]ProductType{
		"installment": INSTALLMENT,
		"loan":        LOAN,
	}
)

func (p ProductType) String() string {
	return productType[p]
}

func ValidateProductType(productType string) (string, error) {
	val, ok := validateProductType[productType]
	if !ok {
		return "", ErrUnknownProductType
	}
	return val.String(), nil
}
