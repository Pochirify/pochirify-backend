package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
)

var (
	errCreateOrder = errors.New("pochirify-backend-internal-usecase-order: failed to create order")
)

func (a App) PayPayTransactionEvent(ctx context.Context) error {
	return nil
}

type CreateOrderInput struct {
	ProductID     string
	Quantity      int
	PaymentMethod model.PaymentMethod
	redirectURL   *string
	UserID        *string
	PhoneNumber   string

	AddressID     *string
	EmailAddress  string
	ZipCode       int
	Prefecture    string
	City          string
	StreetAddress string
	Building      *string
	LastName      string
	FirstName     string
}

type CreateOrderOutput struct {
	OrderID     string
	TotalPrice  int
	OrderOutput *OrderUnion
}

type OrderUnion struct {
	// return only one from
	CreditCardOrder *payment.CreditCardOrder
	PayPayOrder     *payment.PayPayOrder
}

// TODO: どこかバグってるかもしれない。テスト作ってあとで調べる
func (a App) CreateOrder(ctx context.Context, input *CreateOrderInput) (*CreateOrderOutput, error) {
	var user *model.User
	var err error
	if input.UserID != nil {
		user, err = a.UserRepo.Find(ctx, *input.UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
		}
	} else {
		user, err = model.NewUser(input.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
		}
	}

	var userAddress *model.UserAddress
	if input.AddressID != nil {
		userAddress, err = a.UserRepo.FindUserAddress(ctx, *input.AddressID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
		}
	} else {
		userAddress, err = model.NewUserAddress(
			user.ID,
			input.EmailAddress,
			input.ZipCode,
			input.Prefecture,
			input.City,
			input.StreetAddress,
			input.Building,
			input.LastName,
			input.FirstName,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
		}
	}

	// NOTE: 価格はここで確定する
	var totalPrice int
	if product, err := a.ProductRepo.Find(ctx, input.ProductID); err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
	} else {
		totalPrice = product.GetTotalPrice(input.Quantity)
	}

	order := model.NewOrder(
		user,
		userAddress,
		input.PaymentMethod,
		input.ProductID,
		totalPrice,
	)
	orderOutput, err := a.getOrderOutput(ctx, input.PaymentMethod, order.ID, totalPrice, input.redirectURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
	}

	err = a.Tx.Transaction(ctx, func(ctx context.Context) error {
		// TODO: lockの順番
		product, tErr := a.ProductRepo.Find(ctx, input.ProductID)
		if tErr != nil {
			return tErr
		}
		product.Bought()
		if tErr := a.ProductRepo.Update(ctx, product); tErr != nil {
			return tErr
		}

		if tErr := a.OrderRepo.Create(ctx, order); tErr != nil {
			return tErr
		}

		if input.UserID == nil {
			if tErr := a.UserRepo.Create(ctx, user); tErr != nil {
				return tErr
			}
		}
		if input.AddressID == nil {
			if tErr := a.UserRepo.CreateUserAddress(ctx, userAddress); tErr != nil {
				return tErr
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
	}

	return &CreateOrderOutput{
		OrderID:     order.ID,
		TotalPrice:  totalPrice,
		OrderOutput: orderOutput,
	}, nil
}

func (a App) getOrderOutput(
	ctx context.Context,
	paymentMethod model.PaymentMethod,
	orderID string,
	price int,
	redirectURL *string,
) (*OrderUnion, error) {
	switch {
	case paymentMethod.IsPayPay():
		if redirectURL == nil {
			return nil, fmt.Errorf(
				"failed to get orderOutput because redirectURL not provided. paymentMethod=%s",
				paymentMethod.String(),
			)
		}
		qr, err := a.paypayClient.CreateOrder(ctx, orderID, price, *redirectURL)
		if err != nil {
			return nil, err
		}

		return &OrderUnion{
			PayPayOrder: qr,
		}, nil
	case paymentMethod.IsCard():
		order, err := a.creditCardClient.CreateOrder(ctx, orderID, price)
		if err != nil {
			return nil, err
		}

		return &OrderUnion{
			CreditCardOrder: order,
		}, nil
	default:
		return nil, fmt.Errorf("failed to get order output. paymentMethod=%s", paymentMethod.String())
	}
}

// type CreatePaypayQRCodeInput struct {
// 	EmailAddress string
// 	PhoneNumber  string
// 	Zip          string
// 	Prefecture   string
// 	AddressOne   string
// 	AddressTwo   *string

// 	Amount           int
// 	OrderDescription string
// }

// type CreatePaypayQRCodeOutput struct {
// 	QRCode *model.PaypayQRCode
// }

// func (a App) CreatePaypayQRCode(ctx context.Context, input *CreatePaypayQRCodeInput) (*CreatePaypayQRCodeOutput, error) {
// 	user, err := model.NewUser(
// 		input.EmailAddress,
// 		input.PhoneNumber,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
// 	}

// 	userAddress, err := model.NewUserAddress(
// 		user.ID,
// 		input.Zip,
// 		input.Prefecture,
// 		input.AddressOne,
// 		input.AddressTwo,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
// 	}

// 	order := model.NewOrder(user, userAddress, input.Amount, input.OrderDescription)
// 	qrCode, err := a.paypayClient.CreateQRCode(ctx, order.ID, order.OrderDescription)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
// 	}

// 	err = a.Tx.Transaction(ctx, func(ctx context.Context) error {
// 		if tErr := a.UserRepo.Upsert(ctx, user); tErr != nil {
// 			return tErr
// 		}
// 		if tErr := a.UserRepo.CreateUserAddress(ctx, userAddress); tErr != nil {
// 			return tErr
// 		}
// 		if tErr := a.OrderRepo.Create(ctx, order); tErr != nil {
// 			return tErr
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreatePaypayQRCode)
// 	}

// 	return &CreatePaypayQRCodeOutput{QRCode: qrCode}, nil
// }
