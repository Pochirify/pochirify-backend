package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
)

var (
	errCreateOrder                      = errors.New("pochirify-backend-internal-usecase-order: failed to create order")
	ErrCompleteOrderOrderIsNotCompleted = errors.New("pochirify-backend-internal-usecase-order: failed to complete order because order is not completed")
)

func (a App) PayPayTransactionEvent(ctx context.Context) error {
	return nil
}

type CreateOrderInput struct {
	ProductVariantID uint
	Quantity         uint
	PaymentMethod    model.PaymentMethod
	RedirectURL      *string // for paypay

	PhoneNumber   string
	EmailAddress  string
	ZipCode       uint
	Prefecture    string
	City          string
	StreetAddress string
	Building      *string
	LastName      string
	FirstName     string
}

type CreateOrderOutput struct {
	// TODO: add other
	OrderID        string
	TotalPrice     uint
	OrderOutput    *OrderUnion
}

type OrderUnion struct {
	// return only one from
	CreditCardOrder *payment.CreditCardOrder
	PayPayOrder     *payment.CreatOrderPayload
}

func (a App) CreateOrder(ctx context.Context, input *CreateOrderInput) (*CreateOrderOutput, error) {
	order, err := model.NewOrder(
		model.PaymentMethodPayPay,
		input.ProductVariantID,
		input.Quantity,
	)
	if err != nil {
		return nil, err
	}

	shippingAddress, err := model.NewShippingAddress(
		input.EmailAddress,
		input.ZipCode,
		input.Prefecture,
		input.City,
		input.StreetAddress,
		input.Building,
		input.LastName,
		input.FirstName,
		input.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	payload, err := a.shopifyClient.CreatePendingOrder(ctx, input.Quantity, input.ProductVariantID, shippingAddress)
	// TODO: handle inventory
	if err != nil {
		return nil, fmt.Errorf("shopifyClient.CreatePendingOrder failed: %w", err)
	}

	order.Update(uint(payload.ShopifyOrderID), payload.TotalPrice)

	paypayPayload, err := a.paypayClient.CreateOrder(
		ctx,
		order.ID,
		int(order.GetTotalPrice()),
		*input.RedirectURL,
	)
	if err != nil {
		return nil, fmt.Errorf("paypayClient.CreateOrder failed: %w", err)
	}

	if err = a.OrderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	return &CreateOrderOutput{
		OrderID:        order.ID,
		TotalPrice:     order.GetTotalPrice(),
		OrderOutput: &OrderUnion{
			PayPayOrder: &payment.CreatOrderPayload{
				URL: paypayPayload.URL,
			},
		},
	}, nil
}

type CompleteOrderOutput struct {
	ShopifyActivationURL *string
}

func (a App) CompleteOrder(ctx context.Context, id string) (*CompleteOrderOutput, error) {
	order, err := a.OrderRepo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find order: %w", err)
	}

	getOrderPayload, err := a.paypayClient.GetOrder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get paypay order: %w", err)
	}
	if !getOrderPayload.Status.IsStatusCompleted() {
		return nil, fmt.Errorf("paypay order is not completed: %w", ErrCompleteOrderOrderIsNotCompleted)
	}

	shopifyOrderGID, err := a.shopifyClient.MarkOrderAsPaid(ctx, order.ShopifyOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to mark order as paid: %w", err)
	}

	url, err := a.shopifyClient.GetShopifyActivationURL(ctx, shopifyOrderGID)
	if err != nil {
		switch {
		case errors.Is(err, shopify.ErrGetShopifyActivationURLAlreadyActivated):
			break
		default:
			return nil, fmt.Errorf("failed to get shopify activation url: %w", err)
		}
	}

	order.AsPaid()
	if err := a.OrderRepo.Update(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	var activationURl *string
	if url != "" {
		activationURl = &url
	}
	return &CompleteOrderOutput{
		ShopifyActivationURL: activationURl,
	}, nil
}

// // TODO: どこかバグってるかもしれない。テスト作ってあとで調べる
// func (a App) CreateOrder(ctx context.Context, input *CreateOrderInput) (*CreateOrderOutput, error) {
// 	var user *model.User
// 	var err error
// 	if input.UserID != nil {
// 		user, err = a.UserRepo.Find(ctx, *input.UserID)
// 		if err != nil {
// 			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 		}
// 	} else {
// 		user, err = model.NewUser(input.PhoneNumber)
// 		if err != nil {
// 			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 		}
// 	}

// 	// TODO: userAddressをidで指定できる仕様いらなそう
// 	var userAddress *model.UserAddress
// 	if input.AddressID != nil {
// 		userAddress, err = a.UserRepo.FindUserAddress(ctx, *input.AddressID)
// 		if err != nil {
// 			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 		}
// 	} else {
// 		userAddress, err = model.NewUserAddress(
// 			user.ID,
// 			input.EmailAddress,
// 			input.ZipCode,
// 			input.Prefecture,
// 			input.City,
// 			input.StreetAddress,
// 			input.Building,
// 			input.LastName,
// 			input.FirstName,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 		}
// 	}

// 	// NOTE: 価格はここで確定する
// 	var totalPrice int
// 	if product, err := a.ProductRepo.Find(ctx, input.ProductID); err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 	} else {
// 		totalPrice = product.GetTotalPrice(input.Quantity)
// 	}

// 	order := model.NewOrder(
// 		user,
// 		userAddress,
// 		input.PaymentMethod,
// 		input.ProductID,
// 		totalPrice,
// 	)
// 	orderOutput, err := a.getOrderOutput(ctx, input.PaymentMethod, order.ID, totalPrice, input.RedirectURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 	}

// 	err = a.Tx.Transaction(ctx, func(ctx context.Context) error {
// 		// TODO: lockの順番
// 		product, tErr := a.ProductRepo.Find(ctx, input.ProductID)
// 		if tErr != nil {
// 			return tErr
// 		}
// 		product.Bought()
// 		if tErr := a.ProductRepo.Update(ctx, product); tErr != nil {
// 			return tErr
// 		}

// 		if input.UserID == nil {
// 			if tErr := a.UserRepo.Create(ctx, user); tErr != nil {
// 				return tErr
// 			}
// 		}
// 		if input.AddressID == nil {
// 			if tErr := a.UserRepo.CreateUserAddress(ctx, userAddress); tErr != nil {
// 				return tErr
// 			}
// 		}

// 		if tErr := a.OrderRepo.Create(ctx, order); tErr != nil {
// 			return tErr
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errCreateOrder)
// 	}

// 	return &CreateOrderOutput{
// 		OrderID:     order.ID,
// 		TotalPrice:  totalPrice,
// 		OrderOutput: orderOutput,
// 	}, nil
// }

// func (a App) getOrderOutput(
// 	ctx context.Context,
// 	paymentMethod model.PaymentMethod,
// 	orderID string,
// 	price int,
// 	redirectURL *string,
// ) (*OrderUnion, error) {
// 	switch {
// 	case paymentMethod.IsPayPay():
// 		if redirectURL == nil {
// 			return nil, fmt.Errorf(
// 				"failed to get orderOutput because redirectURL not provided. paymentMethod=%s",
// 				paymentMethod.String(),
// 			)
// 		}
// 		qr, err := a.paypayClient.CreateOrder(ctx, orderID, price, *redirectURL)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &OrderUnion{
// 			PayPayOrder: qr,
// 		}, nil
// 	case paymentMethod.IsCard():
// 		order, err := a.creditCardClient.CreateOrder(ctx, orderID, price)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &OrderUnion{
// 			CreditCardOrder: order,
// 		}, nil
// 	default:
// 		return nil, fmt.Errorf("failed to get order output. paymentMethod=%s", paymentMethod.String())
// 	}
// }

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
