package usecase

import (
	"errors"

	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
	"github.com/Pochirify/pochirify-backend/internal/domain/settlement/paypay"
)

var (
	errCreatePaypayQRCode = errors.New("pochirify-backend-internal-usecase-app: failed to create paypayQRCode")
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
