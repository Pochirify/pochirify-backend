package shopify

import (
	"context"
	"fmt"
	"strconv"

	shopify "github.com/Pochirify/pochirify-backend/internal/handler/ec/shopify/gqlgenc"
)

func (c shopifyClient) MarkOrderAsPaid(ctx context.Context, shopifyOrderID uint) (uint, error) {
	res, err := c.graphqlClient.MarkOrderAsPaid(ctx, shopify.OrderMarkAsPaidInput{
		ID: newShopifyOrderGID(shopifyOrderID),
	})
	if err != nil {
		return 0, err
	}

	if len(res.OrderMarkAsPaid.UserErrors) > 0 {
		return 0, fmt.Errorf("failed to mark order as paid: %v", res.OrderMarkAsPaid.UserErrors[0])
	}

	return newShopifyCustomerIDFromGID(res.OrderMarkAsPaid.Order.Customer.ID)
}

func newShopifyOrderGID(shopifyOrderID uint) string {
	return "gid://shopify/Order/" + strconv.Itoa(int(shopifyOrderID))
}

func newShopifyCustomerIDFromGID(gid string) (uint, error) {
	id, err := strconv.Atoi(gid[len("gid://shopify/Customer/"):])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
