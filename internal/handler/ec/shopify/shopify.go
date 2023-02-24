package shopify

import (
	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
)

var (
	_ shopify.ShopifyClient = (*shopifyClient)(nil)
)

type shopifyClient struct {
	restBaseURL string
	accessToken string
}

func NewShopifyClient(accessToken string) shopify.ShopifyClient {
	return &shopifyClient{
		restBaseURL: "https://kounosuke-test.myshopify.com/admin/api/2023-01/",
		accessToken: accessToken,
	}
}
