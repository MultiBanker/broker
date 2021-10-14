package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var ErrIsEmpty = fmt.Errorf("[ERROR] is Empty")

func ValidationIsEmpty(value string) error {
	return errors.Wrap(ErrIsEmpty, value)
}

type Order struct {
	ID                      string            `json:"id" bson:"_id"`
	ReferenceID             string            `json:"referenceId" bson:"reference_id"`
	OrderState              string            `json:"orderState" bson:"order_state"`
	RedirectURL             string            `json:"redirectUrl" bson:"redirect_url"`
	Channel                 string            `json:"channel" bson:"channel"`
	SystemCode              string            `json:"system_code"`
	StateCode               string            `json:"-"`
	MarketCode              string            `json:"-"`
	ProductType             string            `json:"productType" bson:"product_type"`
	PaymentMethod           string            `json:"paymentMethod" bson:"payment_method"`
	IsDelivery              bool              `json:"isDelivery" bson:"is_delivery"`
	TotalCost               string            `json:"totalCost" bson:"total_cost"`
	LoanLength              string            `json:"loanLength" bson:"loan_length"`
	SalesPlace              string            `json:"salesPlace" bson:"sales_place"`
	VerificationId          string            `json:"verificationId"`
	VerificationSMSCode     string            `json:"verificationSmsCode" bson:"verification_sms_code"`
	VerificationSMSDatetime string            `json:"verificationSmsDateTime" bson:"verification_sms_datetime"`
	Customer                Customer          `json:"customer" bson:"customer"`
	Address                 Address           `json:"address" bson:"address"`
	Goods                   []Goods           `json:"goods" bson:"goods"`
	PaymentPartners         []PaymentPartners `json:"paymentPartners"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Customer struct {
	TaxCode    string  `json:"taxCode"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName string  `json:"middleName"`
	Contact    Contact `json:"contact"`
}

func (c Customer) Validate() error {
	var errstrings []string

	if c.TaxCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("taxCode").Error())
	}

	if c.FirstName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("first name").Error())
	}
	if c.LastName == "" {
		errstrings = append(errstrings, ValidationIsEmpty("last name").Error())
	}

	if c.TaxCode == "" {
		errstrings = append(errstrings, ValidationIsEmpty("taxCode").Error())
	}

	if err := c.Contact.Validate(); err != nil {
		errstrings = append(errstrings, c.Contact.Validate().Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

type Address struct {
	Delivery    string `json:"delivery" bson:"delivery"`
	PickupPoint string `json:"pickupPoint" bson:"pickup_point"`
}

func (a Address) Validate() error {
	var errstrings []string
	if a.Delivery == "" {
		errstrings = append(errstrings, ValidationIsEmpty("delivery").Error())
	}

	if a.PickupPoint == "" {
		errstrings = append(errstrings, ValidationIsEmpty("pickup point").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

type Goods struct {
	Category string `json:"category" bson:"category"`
	Brand    string `json:"brand" bson:"brand"`
	Price    string `json:"price" bson:"price"`
	Model    string `json:"model" bson:"model"`
	Image    string `json:"image" bson:"image"`
}

func (a Goods) Validate() error {
	var errstrings []string

	if a.Category == "" {
		errstrings = append(errstrings, ValidationIsEmpty("delivery").Error())
	}

	if a.Brand == "" {
		errstrings = append(errstrings, ValidationIsEmpty("pickup point").Error())
	}

	if a.Price == "" {
		errstrings = append(errstrings, ValidationIsEmpty("price").Error())
	}

	if a.Model == "" {
		errstrings = append(errstrings, ValidationIsEmpty("model").Error())
	}

	if a.Image == "" {
		errstrings = append(errstrings, ValidationIsEmpty("image").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

type Contact struct {
	MobileNumber string `json:"mobileNumber" bson:"mobile_number"`
	Email        string `json:"email" bson:"email"`
}

func (c Contact) Validate() error {
	var errstrings []string

	if c.MobileNumber == "" {
		errstrings = append(errstrings, ValidationIsEmpty("mobile number").Error())
	}

	if c.Email == "" {
		errstrings = append(errstrings, ValidationIsEmpty("email").Error())
	}

	if errstrings != nil {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil

}

type PaymentPartners struct {
	Code string `json:"code"`
}

func (p PaymentPartners) Validate() error {
	if p.Code == "" {
		return ValidationIsEmpty("code partner")
	}
	return nil
}
