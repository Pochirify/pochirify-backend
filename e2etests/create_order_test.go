// go:build e2e
package e2etests

import (
	"context"
	"testing"

	"github.com/Pochirify/pochirify-backend/e2etests/gqlgenc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//NOTE: この商品の在庫がなくなったら、dev環境を使用したこのe2eテストは壊れてしまう。
const existingProductVariantID = 44501604106551

func defaultCreateOrder(
	t *testing.T,
	client gqlgenc.GraphQLClient,
	ctx context.Context,
	emailAddress string,
	redirectURL *string,
) (*gqlgenc.CreateOrder, error) {
	t.Helper()

	var (
		phoneNumber = "08011112222"
	)

	return client.CreateOrder(
		ctx,
		gqlgenc.CreateOrderInput{
			ProductVariantID: existingProductVariantID,
			UnitPrice:        10,
			Quantity:         2,
			PaymentMethod:    gqlgenc.PaymentMethodPaypay,
			RedirectURL:      redirectURL,
			PhoneNumber:      phoneNumber,
			EmailAddress:     emailAddress,
			ZipCode:          1300013,
			Prefecture:       "東京都",
			City:             "墨田区",
			StreetAddress:    "錦糸1-2",
			Building:         nil,
			LastName:         "山田",
			FirstName:        "太郎",
		},
	)
}

func TestCreateOrder_Normal(t *testing.T) {
	client := newClient(t)
	shopifyClient := newShopifyClient(t)
	ctx := context.Background()
	redirectURL := "https://example.com"

	t.Run("create a order using paypay", func(t *testing.T) {
		emailAddress := uuid.NewString() + "@example.com"

		// run
		res, err := defaultCreateOrder(t, client, ctx, emailAddress, &redirectURL)
		require.NoError(t, err)
		require.NotNil(t, res.CreateOrder.Order)
		assert.NotEqual(t, "", res.CreateOrder.Order.OrderID)
		assert.Equal(t, 20, res.CreateOrder.Order.TotalPrice)
		require.NotNil(t, res.CreateOrder.Order.OrderResult)
		require.NotNil(t, res.CreateOrder.Order.OrderResult.PaypayOrderResult)
		assert.NotEqual(t, "", res.CreateOrder.Order.OrderResult.PaypayOrderResult.URL)

		// check record
		order, err := repositories.OrderRepo.Find(ctx, res.CreateOrder.Order.OrderID)
		require.NoError(t, err)

		// check shopify data
		shopifyOrderGID := newShopifyOrderGID(order.ShopifyOrderID)
		orderRes, err := shopifyClient.GetOrder(ctx, shopifyOrderGID)
		require.NoError(t, err)
		require.NotNil(t, orderRes.Order)
		require.NotEqual(t, "", orderRes.Order.ID)
		assert.Equal(t, emailAddress, *orderRes.Order.Email)
		assert.Equal(t, "PENDING", orderRes.Order.DisplayFinancialStatus.String())

		// check papay data
		paypayOrder, err := paypayClient.GetOrder(ctx, order.ID)
		require.NoError(t, err)
		assert.NotEqual(t, "", paypayOrder.Status)
		assert.Equal(t, "CREATED", paypayOrder.Status.String())
	})

	t.Run("create a order using credit card", func(t *testing.T) {})

	t.Run("create a order, but total price has been changed by shopify", func(t *testing.T) {})

	// TODO:
	t.Run("failed to create a order, because inventory is empty", func(t *testing.T) {})
}
