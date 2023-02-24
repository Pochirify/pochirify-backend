// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgenc

import (
	"fmt"
	"io"
	"strconv"
)

type OrderResult interface {
	IsOrderResult()
}

type AllActiveVariantGroupIDs struct {
	Ids []string `json:"ids"`
}

type DeliveryTimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Product struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Price    int      `json:"price"`
	Contents []string `json:"contents"`
	ImageURL string   `json:"imageURL"`
}

type VariantGroup struct {
	ID                  string            `json:"id"`
	Title               string            `json:"title"`
	ImageURLs           []string          `json:"imageURLs"`
	DeliveryTimeRange   DeliveryTimeRange `json:"deliveryTimeRange"`
	FaqImageURL         WebpPngImageURL   `json:"faqImageURL"`
	DescriptionImageURL WebpPngImageURL   `json:"descriptionImageURL"`
	BadgeImageURL       string            `json:"badgeImageURL"`
}

type VariantGroupDetail struct {
	VariantGroup VariantGroup `json:"variantGroup"`
	Variants     []*Product   `json:"variants"`
}

type WebpPngImageURL struct {
	WebpURL string `json:"webpURL"`
	PngURL  string `json:"pngURL"`
}

type CreateOrderInput struct {
	ProductVariantID string        `json:"productVariantID"`
	UnitPrice        int           `json:"unitPrice"`
	Quantity         int           `json:"quantity"`
	PaymentMethod    PaymentMethod `json:"paymentMethod"`
	RedirectURL      *string       `json:"redirectURL,omitempty"`
	PhoneNumber      string        `json:"phoneNumber"`
	EmailAddress     string        `json:"emailAddress"`
	ZipCode          int           `json:"zipCode"`
	Prefecture       string        `json:"prefecture"`
	City             string        `json:"city"`
	StreetAddress    string        `json:"streetAddress"`
	Building         *string       `json:"building,omitempty"`
	LastName         string        `json:"lastName"`
	FirstName        string        `json:"firstName"`
}

type CreateOrderPayload struct {
	OrderID     string      `json:"orderID"`
	TotalPrice  int         `json:"totalPrice"`
	OrderResult OrderResult `json:"orderResult"`
}

type CreditCardResult struct {
	CardOrderID string `json:"cardOrderID"`
	AccessID    string `json:"accessID"`
}

func (CreditCardResult) IsOrderResult() {}

type PaypayOrderResult struct {
	URL string `json:"url"`
}

func (PaypayOrderResult) IsOrderResult() {}

type PaymentMethod string

const (
	PaymentMethodCard      PaymentMethod = "CARD"
	PaymentMethodPaypay    PaymentMethod = "PAYPAY"
	PaymentMethodApplePay  PaymentMethod = "APPLE_PAY"
	PaymentMethodGooglePay PaymentMethod = "GOOGLE_PAY"
)

var AllPaymentMethod = []PaymentMethod{
	PaymentMethodCard,
	PaymentMethodPaypay,
	PaymentMethodApplePay,
	PaymentMethodGooglePay,
}

func (e PaymentMethod) IsValid() bool {
	switch e {
	case PaymentMethodCard, PaymentMethodPaypay, PaymentMethodApplePay, PaymentMethodGooglePay:
		return true
	}
	return false
}

func (e PaymentMethod) String() string {
	return string(e)
}

func (e *PaymentMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentMethod", str)
	}
	return nil
}

func (e PaymentMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
