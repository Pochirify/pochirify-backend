// go:build e2e
package e2etests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var activatedShopifyCustomerEmail = "dfca4b67-ede4-47c8-bae8-cd40c98e202e@example.com"

func TestCompleteOrder_Normal(t *testing.T) {
	client := newClient(t)
	shopifyClient := newShopifyClient(t)
	ctx := context.Background()

	redirectURL := "https://example.com"

	t.Run("complete a order", func(t *testing.T) {
		emailAddress := uuid.NewString() + "@example.com"
		res, err := defaultCreateOrder(t, client, ctx, emailAddress, &redirectURL)
		require.NoError(t, err)
		require.NotEqual(t, "", res.CreateOrder.OrderID)

		resp, err := client.CompleteOrder(ctx, res.CreateOrder.OrderID)
		require.NoError(t, err)
		assert.NotNil(t, resp.CompleteOrder.ShopifyActivationURL)
		assert.NotEqual(t, "", resp.CompleteOrder.ShopifyActivationURL)

		order, err := repositories.OrderRepo.Find(ctx, res.CreateOrder.OrderID)
		require.NoError(t, err)
		assert.Equal(t, "completed", order.Status.String())

		shopifyOrderGID := newShopifyOrderGID(order.ShopifyOrderID)
		orderRes, err := shopifyClient.GetOrder(ctx, shopifyOrderGID)
		require.NoError(t, err)
		require.NotNil(t, orderRes.Order)
		require.NotEqual(t, "", orderRes.Order.ID)
		assert.Equal(t, "PAID", orderRes.Order.DisplayFinancialStatus.String())
	})

	t.Run("complete a order without activation url, because already activated", func(t *testing.T) {
		res, err := defaultCreateOrder(t, client, ctx, activatedShopifyCustomerEmail, &redirectURL)
		require.NoError(t, err)
		require.NotEqual(t, "", res.CreateOrder.OrderID)

		resp, err := client.CompleteOrder(ctx, res.CreateOrder.OrderID)
		require.NoError(t, err)
		assert.Nil(t, resp.CompleteOrder.ShopifyActivationURL)

		// TODO: define verifying func
		order, err := repositories.OrderRepo.Find(ctx, res.CreateOrder.OrderID)
		require.NoError(t, err)
		assert.Equal(t, "completed", order.Status.String())

		shopifyOrderGID := newShopifyOrderGID(order.ShopifyOrderID)
		orderRes, err := shopifyClient.GetOrder(ctx, shopifyOrderGID)
		require.NoError(t, err)
		require.NotNil(t, orderRes.Order)
		require.NotEqual(t, "", orderRes.Order.ID)
		assert.Equal(t, "PAID", orderRes.Order.DisplayFinancialStatus.String())
	})

	// TODO:
	t.Run("failed to complete a order, because paypay checkout has not been completed", func(t *testing.T) {})
}
