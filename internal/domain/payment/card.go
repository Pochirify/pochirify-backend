package payment

import "context"

type CreditCardClient interface {
	CreateOrder(ctx context.Context, orderID string, amount int) (*CreditCardOrder, error)
}

type CreditCardOrder struct {
	CardOrderID string
	AccessID    string
}
