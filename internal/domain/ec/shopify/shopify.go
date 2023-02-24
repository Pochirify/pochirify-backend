package shopify

import (
	"context"
	"errors"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

var ErrCreateOrderUnableToReserveInventory = errors.New("pochirify-backend-internal-domain-model-shopify: failed to create order because unable to reserve inventory")

type CreateOrderPayload struct {
	ShopifyOrderID int
	IsPriceChanged bool
	TotalPrice     uint
}

type ShopifyClient interface {
	CreateOrder(
		ctx context.Context,
		order *model.Order,
		shippingAddress *model.ShippingAddress,
	) (*CreateOrderPayload, error)
}
