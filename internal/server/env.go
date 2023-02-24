package server

import (
	"github.com/kelseyhightower/envconfig"
)

type environments struct {
	Port              uint   `envconfig:"PORT" required:"true"`
	GCPProjectID      string `envconfig:"GCP_PROJECT_ID" required:"true"`
	SpannerInstanceID string `envconfig:"SPANNER_INSTANCE_ID" required:"true"`
	SpannerDatabaseID string `envconfig:"SPANNER_DATABASE_ID" required:"true"`

	// payment
	FincodeApiKey      string `envconfig:"FINCODE_API_KEY" required:"true"`
	FincodeBaseURL     string `envconfig:"FINCODE_BASE_URL" required:"true"`
	IsPayPayProduction bool   `envconfig:"IS_PAYPAY_PRODUCTION" required:"true"`
	PayPayApiKeyID     string `envconfig:"PAYPAY_API_KEY_ID" required:"true"`
	PayPayApiSecret    string `envconfig:"PAYPAY_API_SECRET" required:"true"`
	PayPayMerchantID   string `envconfig:"PAYPAY_MERCHANT_ID" required:"true"`
	// TODO: いらん
	PayPayRedirectURL string `envconfig:"PAYPAY_REDIRECT_URL" required:"true"`

	ShopifyAdminAccessToken string `envconfig:"SHOPIFY_ADMIN_ACCESS_TOKEN" required:"true"`
}

func getEnvironments() (*environments, error) {
	var e environments
	if err := envconfig.Process("", &e); err != nil {
		return nil, err
	}

	return &e, nil
}
