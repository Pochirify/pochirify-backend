// go:build e2e
package e2etests

import (
	"context"
	"testing"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var activatedShopifyCustomerEmail = "dfca4b67-ede4-47c8-bae8-cd40c98e202e@example.com"

func TestCompleteOrder_Normal(t *testing.T) {
	client := newClient(t)
	shopifyClient := newShopifyClient(t)
	ctx := context.Background()

	redirectURL := "https://example.com"

	verifyData := func(t *testing.T, orderID string, status model.PaymentStatus, shopifyFinancialStatus string) {
		t.Helper()
		order, err := repositories.OrderRepo.Find(ctx, orderID)
		require.NoError(t, err)
		assert.Equal(t, status.String(), order.Status.String())

		shopifyOrderGID := newShopifyOrderGID(order.ShopifyOrderID)
		orderRes, err := shopifyClient.GetOrder(ctx, shopifyOrderGID)
		require.NoError(t, err)
		require.NotNil(t, orderRes.Order)
		require.NotEqual(t, "", orderRes.Order.ID)
		assert.Equal(t, shopifyFinancialStatus, orderRes.Order.DisplayFinancialStatus.String())
	}

	// paypay決済をapiで任意のタイミング行えないため、e2eで正常系を再現はできない
	// なので、usecase層でmockを挿しつつUTを書くべき。

	// t.Run("complete a order", func(t *testing.T) {
	// 	emailAddress := uuid.NewString() + "@example.com"
	// 	res, err := defaultCreateOrder(t, client, ctx, emailAddress, &redirectURL)
	// 	require.NoError(t, err)
	// 	require.NotNil(t, res.CreateOrder.Order)
	// 	assert.NotEqual(t, "", res.CreateOrder.Order.OrderID)

	// 	// complete order with paypay here

	// 	paypayOrder, err := paypayClient.GetOrder(ctx, res.CreateOrder.Order.OrderID)
	// 	require.NoError(t, err)
	// 	assert.Equal(t, payment.PayPayOrderCreated, paypayOrder.Status)

	// 	resp, err := client.CompleteOrder(ctx, "55c71860-64bc-447a-918e-e9ce529009cd")
	// 	require.NoError(t, err)
	// 	assert.NotNil(t, resp.CompleteOrder.ShopifyActivationURL)
	// 	assert.NotEqual(t, "", resp.CompleteOrder.ShopifyActivationURL)

	// 	verifyData(t, res.CreateOrder.Order.OrderID, model.PaymentStatusCompleted, "PAID")
	// })

	// t.Run("complete a order without activation url, because already activated", func(t *testing.T) {
	// 	res, err := defaultCreateOrder(t, client, ctx, activatedShopifyCustomerEmail, &redirectURL)
	// 	require.NoError(t, err)
	// 	require.NotNil(t, res.CreateOrder.Order)
	// 	assert.NotEqual(t, "", res.CreateOrder.Order.OrderID)

	// 	// complete order with paypay here

	// 	resp, err := client.CompleteOrder(ctx, res.CreateOrder.Order.OrderID)
	// 	require.NoError(t, err)
	// 	assert.Nil(t, resp.CompleteOrder.ShopifyActivationURL)
	// 	assert.False(t, resp.CompleteOrder.IsNotOrderCompleted)

	// 	verifyData(t, res.CreateOrder.Order.OrderID, model.PaymentStatusCompleted, "PAID")
	// })

	t.Run("failed to complete a order, because paypay checkout has not been completed", func(t *testing.T) {
		res, err := defaultCreateOrder(t, client, ctx, activatedShopifyCustomerEmail, &redirectURL)
		require.NoError(t, err)
		require.NotNil(t, res.CreateOrder.Order)
		assert.NotEqual(t, "", res.CreateOrder.Order.OrderID)

		resp, err := client.CompleteOrder(ctx, res.CreateOrder.Order.OrderID)
		require.NoError(t, err)
		assert.True(t, resp.CompleteOrder.IsNotOrderCompleted)

		// check paypay order
		paypayOrder, err := paypayClient.GetOrder(ctx, res.CreateOrder.Order.OrderID)
		require.NoError(t, err)
		assert.Equal(t, payment.PayPayOrderCreated, paypayOrder.Status)

		verifyData(t, res.CreateOrder.Order.OrderID, model.PaymentStatusPending, "PENDING")
	})
}
