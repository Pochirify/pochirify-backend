package payment

import (
	"context"
)

type PaypayClient interface {
	CreateOrder(ctx context.Context, orderID string, price int, redirectURL string) (*PayPayOrder, error)
}

type PayPayOrder struct {
	URL string
}
