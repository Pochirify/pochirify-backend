package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/mythrnr/paypayopa-sdk-go"
)

var errCreatePayPayOrder = errors.New("pochirify-backend-internal-handler-payment-paypay: failed to create order")

type Environment int

const (
	_ Environment = iota
	EnvProduction
	EnvSandbox
)

type paypayClient struct {
	isProduction bool
	apiKeyID     string
	apiKeySecret string
	merchantID   string
	redirectURL  string
}

func NewPaypayClient(
	isProduction bool,
	apiKeyID,
	apiKeySecret,
	merchantIDcreditCardClient,
	redirectURL string,
) payment.PaypayClient {
	return &paypayClient{
		isProduction: isProduction,
		apiKeyID:     apiKeyID,
		apiKeySecret: apiKeySecret,
		merchantID:   merchantIDcreditCardClient,
		redirectURL:  redirectURL,
	}
}

func (r *paypayClient) CreateOrder(ctx context.Context, orderID string, price int) (*payment.PayPayOrder, error) {
	cred := paypayopa.NewCredentials(
		r.getEnv(),
		r.apiKeyID,
		r.apiKeySecret,
		r.merchantID,
	)
	wp := paypayopa.NewWebPayment(cred)

	res, info, err := wp.CreateQRCode(ctx, &paypayopa.CreateQRCodePayload{
		MerchantPaymentID: orderID,
		Amount: &paypayopa.MoneyAmount{
			Amount:   price,
			Currency: paypayopa.CurrencyJPY,
		},
		CodeType:     paypayopa.CodeTypeOrderQR,
		RedirectURL:  r.redirectURL,
		RedirectType: paypayopa.RedirectTypeWebLink,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreatePayPayOrder)
	}
	if !info.Success() {
		return nil, fmt.Errorf("%v: %w", info, errCreatePayPayOrder)
	}

	return &payment.PayPayOrder{
		URL: res.URL,
	}, nil
}

func (r paypayClient) getEnv() paypayopa.Environment {
	if r.isProduction {
		return paypayopa.EnvProduction
	}
	return paypayopa.EnvSandbox
}
