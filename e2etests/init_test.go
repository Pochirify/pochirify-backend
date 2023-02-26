// go:build e2e
package e2etests

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	appspanner "github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger/json"
	hpayment "github.com/Pochirify/pochirify-backend/internal/handler/payment"
	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/Pochirify/pochirify-backend/e2etests/gqlgenc"
	shopify "github.com/Pochirify/pochirify-backend/e2etests/shopify/gqlgenc"

	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

func newClient(_ *testing.T) gqlgenc.GraphQLClient {
	return gqlgenc.NewClient(
		http.DefaultClient,
		"http://localhost:"+port+"/api/query",
	)
}

func newShopifyClient(_ *testing.T) shopify.ShopifyClient {
	return shopify.NewClient(
		http.DefaultClient,
		"https://kounosuke-test.myshopify.com/admin/api/2023-01/graphql.json",
		func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
			req.Header.Add("X-Shopify-Access-Token", shopifyAccessToken)
			return next(ctx, req, gqlInfo, res)
		},
	)
}

func initRepositories() repository.Repositories {
	client, err := spanner.NewClient(
		context.Background(),
		&appspanner.ClientConfig{
			ProjectID:  projectID,
			InstanceID: instanceID,
			DatabaseID: databaseID,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create spanner client: %s", err.Error()))
	}

	return appspanner.InitRepositories(appspanner.NewSpanner(client, newLoggerFactory()))
}

func newPaypayClient(
	isPayPayProduction,
	paypayApiKeyID,
	paypayApiSecret,
	paypayMerchantID string,
) payment.PaypayClient {
	isProduction := false
	if isPayPayProduction == "true" {
		isProduction = true
	}
	if paypayApiKeyID == "" {
		log.Println("paypay api key id not provided")
	}
	if paypayApiSecret == "" {
		log.Panicln("paypay api secret not provided")
	}
	if paypayMerchantID == "" {
		log.Println("paypay merchant id not provided")
	}
	return hpayment.NewPaypayClient(
		isProduction,
		paypayApiKeyID,
		paypayApiSecret,
		paypayMerchantID,
	)
}

func newLoggerFactory() logger.Factory {
	l, err := json.NewLogger()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize new logger: %s", err.Error()))
	}
	return func(ctx context.Context) logger.Logger {
		return l
	}
}

func newShopifyOrderGID(shopifyOrderID uint) string {
	return "gid://shopify/Order/" + strconv.Itoa(int(shopifyOrderID))
}
