package server

import (
	"context"
	"fmt"
	"os"

	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	"github.com/Pochirify/pochirify-backend/internal/handler/ec/shopify"
	"github.com/Pochirify/pochirify-backend/internal/handler/http"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger/json"
	"github.com/Pochirify/pochirify-backend/internal/handler/payment"
)

func Run() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	env, err := getEnvironments()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	l, err := json.NewLogger()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	spannerClient, err := spanner.NewClient(
		ctx,
		&spanner.ClientConfig{
			ProjectID:  env.GCPProjectID,
			InstanceID: env.SpannerInstanceID,
			DatabaseID: env.SpannerDatabaseID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create spanner client: %w", err)
	}

	config := &http.Config{
		Port:    env.Port,
		Logger:  l,
		Spanner: spannerClient,
		PayPayClient: payment.NewPaypayClient(
			env.IsPayPayProduction,
			env.PayPayApiKeyID,
			env.PayPayApiSecret,
			env.PayPayMerchantID,
		),
		CreditCardClient: payment.NewCreditCardClient(env.FincodeApiKey, env.FincodeBaseURL),
		ShopifyClient:    shopify.NewShopifyClient(env.ShopifyAdminAccessToken),
	}

	// TODO: graceful shutdown
	return http.NewServer(ctx, config).Start()
}
