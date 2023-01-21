package usecase

import (
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

type App struct {
	paypayClient     payment.PaypayClient
	creditCardClient payment.CreditCardClient
	repository.Repositories
}

type Config struct {
	PaypayClient     payment.PaypayClient
	CreditCardClient payment.CreditCardClient
	Repositories     repository.Repositories
}

func NewApp(c *Config) App {
	return App{
		paypayClient:     c.PaypayClient,
		creditCardClient: c.CreditCardClient,
		Repositories:     c.Repositories,
	}
}
