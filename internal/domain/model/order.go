package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	PaymentStatus int
	PaymentMethod int
)

const (
	PaymentStatusUnknown PaymentStatus = iota
	PaymentStatusCreated
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
	ID            string
	UserID        string
	UserAddressID string
	Status        PaymentStatus
	PaymentMethod PaymentMethod
	ProductID     string
	Price         int
	CreateTime    time.Time
	UpdateTime    time.Time
}

func NewOrder(
	user *User,
	userAddress *UserAddress,
	paymentMethod PaymentMethod,
	productID string,
	price int,
) *Order {
	return &Order{
		ID:            uuid.NewString(),
		UserID:        user.ID,
		UserAddressID: userAddress.ID,
		Status:        PaymentStatusCreated,
		PaymentMethod: paymentMethod,
		ProductID:     productID,
		Price:         price,
	}
}

func (m PaymentMethod) IsPayPay() bool {
	return m == PaymentMethodPayPay
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

func (s PaymentStatus) String() string {
	switch s {
	case PaymentStatusCreated:
		return "created"
	case PaymentStatusCompleted:
		return "completed"
	case PaymentStatusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}
