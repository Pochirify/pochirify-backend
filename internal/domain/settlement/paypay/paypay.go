package paypay

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

type PaypayClient interface {
	CreateQRCode(ctx context.Context, merchantPaymentId, orderDescription string) (*model.PaypayQRCode, error)
}

type paypayClient struct {
}

func NewPaypayClient() PaypayClient {
	return &paypayClient{}
}

func (r *paypayClient) CreateQRCode(ctx context.Context, merchantPaymentId, orderDescription string) (*model.PaypayQRCode, error) {
	return &model.PaypayQRCode{
		QRCodeUrl:      "url",
		QRCodeDeepLink: "deepLink",
	}, nil
}
