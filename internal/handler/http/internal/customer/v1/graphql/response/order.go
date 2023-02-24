package response

import (
	"errors"

	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql"
	"github.com/Pochirify/pochirify-backend/internal/usecase"
)

func NewCreateOrderPayload(output *usecase.CreateOrderOutput) (*graphql.CreateOrderPayload, error) {
	orderResult, err := getOrderResult(output.OrderOutput)
	if err != nil {
		return nil, err
	}

	return &graphql.CreateOrderPayload{
		OrderID:     output.OrderID,
		TotalPrice:  int(output.TotalPrice),
		OrderResult: orderResult,
	}, nil
}

func getOrderResult(orderUnion *usecase.OrderUnion) (graphql.OrderResult, error) {
	switch {
	case orderUnion.CreditCardOrder != nil &&
		orderUnion.PayPayOrder == nil:
		return graphql.CreditCardResult{
			CardOrderID: orderUnion.CreditCardOrder.CardOrderID,
			AccessID:    orderUnion.CreditCardOrder.AccessID,
		}, nil
	case orderUnion.PayPayOrder != nil &&
		orderUnion.CreditCardOrder == nil:
		return graphql.PaypayOrderResult{
			URL: orderUnion.PayPayOrder.URL,
		}, nil
	default:
		return nil, errors.New("pochirify-backend-handler-http-internal-customer-response: failed to get orderResult")
	}
}
