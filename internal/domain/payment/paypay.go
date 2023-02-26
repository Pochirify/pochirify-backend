package payment

import (
	"context"
)

type PayPayOrderStatus int

const (
	_ PayPayOrderStatus = iota
	PayPayOrderCreated
	PayPayOrderCompleted
	PayPayOrderCanceled
)

type PaypayClient interface {
	GetOrder(ctx context.Context, orderID string) (*GetOrderPayload, error)
	CreateOrder(ctx context.Context, orderID string, totalPrice int, redirectURL string) (*CreatOrderPayload, error)
}

type GetOrderPayload struct {
	Status PayPayOrderStatus
}

type CreatOrderPayload struct {
	URL string
}

func GetPayPayOrderStatus(status string) PayPayOrderStatus {
	switch status {
	case "CREATED":
		return PayPayOrderCreated
	case "COMPLETED":
		return PayPayOrderCompleted
	case "CANCELED":
		return PayPayOrderCanceled
	default:
		return 0
	}
}

func (s PayPayOrderStatus) String() string {
	switch s {
	case PayPayOrderCreated:
		return "CREATED"
	case PayPayOrderCompleted:
		return "COMPLETED"
	case PayPayOrderCanceled:
		return "CANCELED"
	default:
		return "unknown"
	}
}

func (s PayPayOrderStatus) IsStatusCompleted() bool {
	return s == PayPayOrderCompleted
}
