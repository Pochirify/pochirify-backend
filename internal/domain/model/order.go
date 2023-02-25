package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	PaymentStatus int
	PaymentMethod int
)

const (
	PaymentStatusUnknown PaymentStatus = iota
	PaymentStatusPending
	PaymentStatusCompleted
	PaymentStatusCanceled
)

const (
	PaymentMethodUnknown PaymentMethod = iota
	PaymentMethodCard
	PaymentMethodPayPay
	PaymentMethodApplePay
	PaymentMethodGooglePay
)

type Order struct {
	ID             string // shopify order id
	ShopifyOrderID uint
	// UserID        string
	// UserAddressID string
	Status           PaymentStatus
	PaymentMethod    PaymentMethod
	ProductVariantID uint
	UnitPrice        uint
	Quantity         uint

	CreateTime time.Time
	UpdateTime time.Time
}

func NewOrder(
	paymentMethod PaymentMethod,
	productVariantID,
	unitPrice uint,
	quantity uint,
) (*Order, error) {
	if quantity == 0 {
		return nil, errors.New("quantity must be greater than 0")
	}
	return &Order{
		ID:               uuid.NewString(),
		Status:           PaymentStatusPending,
		PaymentMethod:    paymentMethod,
		ProductVariantID: productVariantID,
		UnitPrice:        unitPrice,
		Quantity:         quantity,
	}, nil
}

func (o *Order) Update(shopifyOrderID, totalPrice uint) {
	o.ShopifyOrderID = shopifyOrderID
	o.UnitPrice = totalPrice / o.Quantity
}

func (o Order) GetTotalPrice() uint {
	return o.UnitPrice * o.Quantity
}

func (m PaymentMethod) IsPayPay() bool {
	return m == PaymentMethodPayPay
}

func (m PaymentMethod) IsCard() bool {
	return m == PaymentMethodCard
}

func (m PaymentMethod) String() string {
	switch m {
	case PaymentMethodCard:
		return "card"
	case PaymentMethodPayPay:
		return "paypay"
	case PaymentMethodApplePay:
		return "apple_pay"
	case PaymentMethodGooglePay:
		return "google_pay"
	default:
		return "unknown"
	}
}

func GetPaymentMethod(pm string) PaymentMethod {
	switch pm {
	case "card":
		return PaymentMethodCard
	case "paypay":
		return PaymentMethodPayPay
	case "apple_pay":
		return PaymentMethodApplePay
	case "google_pay":
		return PaymentMethodGooglePay
	default:
		return PaymentMethodUnknown
	}
}

func (s PaymentStatus) String() string {
	switch s {
	case PaymentStatusPending:
		return "pending"
	case PaymentStatusCompleted:
		return "completed"
	case PaymentStatusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

func GetOrderPaymentStatus(ps string) PaymentStatus {
	switch ps {
	case "pending":
		return PaymentStatusPending
	case "completed":
		return PaymentStatusCompleted
	case "canceled":
		return PaymentStatusCanceled
	default:
		return PaymentStatusUnknown
	}
}

func (o *Order) AsPaid() {
	o.Status = PaymentStatusCompleted
}
