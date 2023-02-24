package usecase

import (
	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

type App struct {
	paypayClient     payment.PaypayClient
	creditCardClient payment.CreditCardClient
	shopifyClient    shopify.ShopifyClient
	repository.Repositories
}

type Config struct {
	PaypayClient     payment.PaypayClient
	CreditCardClient payment.CreditCardClient
	shopify.ShopifyClient
	Repositories repository.Repositories
}

// TODO: seperate usecase and follow ISP law
func NewApp(c *Config) App {
	return App{
		paypayClient:     c.PaypayClient,
		creditCardClient: c.CreditCardClient,
		shopifyClient:    c.ShopifyClient,
		Repositories:     c.Repositories,
	}
}
