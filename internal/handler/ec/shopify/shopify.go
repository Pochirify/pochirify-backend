package shopify

import (
	"context"
	"net/http"

	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
	gqlgenc "github.com/Pochirify/pochirify-backend/internal/handler/ec/shopify/gqlgenc"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"github.com/Yamashou/gqlgenc/clientv2"
)

var (
	_              shopify.ShopifyClient = (*shopifyClient)(nil)
	graphqlBaseURl                       = "https://kounosuke-test.myshopify.com/admin/api/2023-01/graphql.json"
)

type shopifyClient struct {
	restBaseURL   string
	accessToken   string
	graphqlClient gqlgenc.ShopifyAdminClient
	logger logger.Factory
}

func NewShopifyClient(accessToken string, logger logger.Factory) shopify.ShopifyClient {
	return &shopifyClient{
		restBaseURL: "https://kounosuke-test.myshopify.com/admin/api/2023-01/",
		accessToken: accessToken,
		graphqlClient: gqlgenc.NewClient(
			http.DefaultClient,
			graphqlBaseURl,
			func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
				req.Header.Add("X-Shopify-Access-Token", accessToken)
				return next(ctx, req, gqlInfo, res)
			},
		),
		logger: logger,
	}
}
