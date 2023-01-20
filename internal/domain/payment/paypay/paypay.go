package paypay

import (
	"context"
)

type PaypayClient interface {
	CreateQRCode(ctx context.Context, merchantPaymentId, orderDescription string) (*PaypayQRCode, error)
}

type paypayClient struct {
}

func NewPaypayClient() PaypayClient {
	return &paypayClient{}
}

func (r *paypayClient) CreateQRCode(ctx context.Context, merchantPaymentId, orderDescription string) (*PaypayQRCode, error) {
	return &PaypayQRCode{
		QRCodeUrl: "url",
	}, nil
}

type PaypayQRCode struct {
	QRCodeUrl string
}
