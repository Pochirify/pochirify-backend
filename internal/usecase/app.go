package usecase

import (
	"github.com/Pochirify/pochirify-backend/internal/domain/payment/paypay"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

type App struct {
	paypayClient paypay.PaypayClient
	repository.Repositories
}

type Config struct {
	PaypayClient paypay.PaypayClient
	Repositories repository.Repositories
}

func NewApp(c *Config) App {
	return App{
		paypayClient: c.PaypayClient,
		Repositories: c.Repositories,
	}
}
