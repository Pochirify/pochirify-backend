package payment

import (
	"context"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/mythrnr/paypayopa-sdk-go"
)

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
}

func NewPaypayClient(
	isProduction bool,
	apiKeyID,
	apiKeySecret,
	merchantID string,
) payment.PaypayClient {
	return &paypayClient{
		isProduction: isProduction,
		apiKeyID:     apiKeyID,
		apiKeySecret: apiKeySecret,
		merchantID:   merchantID,
	}
}

func (r *paypayClient) GetOrder(ctx context.Context, orderID string) (*payment.GetOrderPayload, error) {
	cred := paypayopa.NewCredentials(
		r.getEnv(),
		r.apiKeyID,
		r.apiKeySecret,
		r.merchantID,
	)
	wp := paypayopa.NewWebPayment(cred)

	res, info, err := wp.GetPaymentDetails(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed GetOrder: %w", err)
	}
	if !info.Success() {
		return nil, fmt.Errorf("failed GetOrder: %v", info)
	}

	return &payment.GetOrderPayload{
		Status: payment.GetPayPayOrderStatus(res.Status),
	}, nil
}

func (r *paypayClient) CreateOrder(ctx context.Context, orderID string, totalPrice int, redirectURL string) (*payment.CreatOrderPayload, error) {
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
			Amount:   totalPrice,
			Currency: paypayopa.CurrencyJPY,
		},
		CodeType:     paypayopa.CodeTypeOrderQR,
		RedirectURL:  redirectURL,
		RedirectType: paypayopa.RedirectTypeWebLink,
	})
	if err != nil {
		return nil, fmt.Errorf("failed CreateOrder: %w", err)
	}
	if !info.Success() {
		return nil, fmt.Errorf("failed CreateOrder: %v", info)
	}

	return &payment.CreatOrderPayload{
		URL: res.URL,
	}, nil
}

func (r paypayClient) getEnv() paypayopa.Environment {
	if r.isProduction {
		return paypayopa.EnvProduction
	}
	return paypayopa.EnvSandbox
}
