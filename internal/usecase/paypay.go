package usecase

import (
	"context"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

func (a App) PayPayTransactionEvent(ctx context.Context) error {
	return nil
}

type CreatePaypayQRCodeInput struct {
	EmailAddress string
	PhoneNumber  string
	Zip          string
	Prefecture   string
	AddressOne   string
	AddressTwo   *string

	Amount           int
	OrderDescription string
}

type CreatePaypayQRCodeOutput struct {
	QRCode *model.PaypayQRCode
}

func (a App) CreatePaypayQRCode(ctx context.Context, input *CreatePaypayQRCodeInput) (*CreatePaypayQRCodeOutput, error) {
	user, err := model.NewUser(
		input.EmailAddress,
		input.PhoneNumber,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
	}

	userAddress, err := model.NewUserAddress(
		user.ID,
		input.Zip,
		input.Prefecture,
		input.AddressOne,
		input.AddressTwo,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
	}

	order := model.NewOrder(user, userAddress, input.Amount, input.OrderDescription)
	qrCode, err := a.paypayClient.CreateQRCode(ctx, order.ID, order.OrderDescription)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
	}

	err = a.Tx.Transaction(ctx, func(ctx context.Context) error {
		if tErr := a.UserRepo.Upsert(ctx, user); tErr != nil {
			return tErr
		}
		if tErr := a.UserRepo.CreateUserAddress(ctx, userAddress); tErr != nil {
			return tErr
		}
		if tErr := a.OrderRepo.Create(ctx, order); tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
	}

	return &CreatePaypayQRCodeOutput{QRCode: qrCode}, nil
}
