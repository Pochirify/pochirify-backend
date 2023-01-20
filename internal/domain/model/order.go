package model

import "github.com/google/uuid"

type PaymentMethod int

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
	PaymentMethod PaymentMethod
	ProductID     string
	Price         int
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
		PaymentMethod: paymentMethod,
		ProductID:     productID,
		Price:         price,
	}
}

func (m PaymentMethod) IsPayPay() bool {
	return m == PaymentMethodPayPay
}
