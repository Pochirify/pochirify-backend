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

const existingProductVariantID = "44501604106551"

func TestCreateOrder_Normal(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	var (
		redirectURL  = "https://example.com"
		phoneNumber  = "08011112222"
		emailAddress = uuid.NewString() + "@example.com"
	)

	t.Run("create a order using paypay", func(t *testing.T) {
		res, err := client.CreateOrder(
			ctx,
			gqlgenc.CreateOrderInput{
				ProductVariantID: existingProductVariantID,
				UnitPrice:        10,
				Quantity:         2,
				PaymentMethod:    gqlgenc.PaymentMethodPaypay,
				RedirectURL:      &redirectURL,
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
		require.NoError(t, err)
		assert.NotEqual(t, "", res.CreateOrder.OrderID)
		assert.Equal(t, 20, res.CreateOrder.TotalPrice)
		require.NotNil(t, res.CreateOrder.OrderResult)
		require.NotNil(t, res.CreateOrder.OrderResult.PaypayOrderResult)
		assert.Equal(t, redirectURL, res.CreateOrder.OrderResult.PaypayOrderResult.URL)

		// TODO: check order record
		// TODO: check paypay order is properly created
	})

	t.Run("create a order using credit card", func(t *testing.T) {})

	t.Run("create a order, but total price has been changed by shopify", func(t *testing.T) {})

	t.Run("failed to create a order, because inventory is empty", func(t *testing.T) {})
}
