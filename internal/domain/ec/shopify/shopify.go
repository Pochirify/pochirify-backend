package shopify

import (
	"context"
	"errors"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

var (
	ErrCreatePendingOrderUnableToReserveInventory = errors.New("pochirify-backend-internal-domain-model-shopify: failed to create order because unable to reserve inventory")
	ErrGetShopifyActivationURLAlreadyActivated    = errors.New("pochirify-backend-internal-domain-model-shopify: failed to get shopify activation url because already activated")
)

type CreatePendingOrderPayload struct {
	ShopifyOrderID int
	TotalPrice     uint
}

type ShopifyClient interface {
	CreatePendingOrder(
		ctx context.Context,
		quantity,
		productVariantID uint,
		shippingAddress *model.ShippingAddress,
	) (*CreatePendingOrderPayload, error)
	// return shopify customer gid
	MarkOrderAsPaid(
		ctx context.Context,
		shopifyOrderID uint,
	) (uint, error)
	GetShopifyActivationURL(ctx context.Context, shopifyCustomerID uint) (string, error)
}
