// go:build e2e
package e2etests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Pochirify/pochirify-backend/e2etests/gqlgenc"
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder_Normal(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	var (
		productID    = uuid.NewString()
		redirectURL  = "https://example.com"
		phoneNumber  = "08011112222"
		emailAddress = uuid.NewString() + "@example.com"
	)

	t.Run("create new user and order", func(t *testing.T) {
		err := repositories.ProductRepo.Create(ctx, &model.Product{
			ID:         productID,
			Title:      uuid.NewString(),
			Price:      10,
			Stock:      100,
			ContentOne: uuid.NewString(),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
		require.NoError(t, err)

		res, err := client.CreateOrder(
			context.Background(),
			gqlgenc.CreateOrderInput{
				ProductID:     productID,
				Quantity:      1,
				PaymentMethod: gqlgenc.PaymentMethodPaypay,
				RedirectURL:   &redirectURL,
				UserID:        nil,
				PhoneNumber:   phoneNumber,
				AddressID:     nil, // TODO:
				EmailAddress:  emailAddress,
				ZipCode:       1830012,
				Prefecture:    "東京都",
				City:          "府中市",
				StreetAddress: "押立町1-33-7",
				Building:      nil, // TODO:
				LastName:      "山田",
				FirstName:     "太郎",
			},
		)
		log.Println(res)
		assert.NoError(t, err)
	})
}
