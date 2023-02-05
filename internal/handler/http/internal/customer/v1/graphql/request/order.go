package request

import (
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql"
	"github.com/Pochirify/pochirify-backend/internal/usecase"
)

func NewCreateOrderInput(i graphql.CreateOrderInput) *usecase.CreateOrderInput {
	return &usecase.CreateOrderInput{
		ProductID:     i.ProductID,
		Quantity:      i.Quantity,
		PaymentMethod: getPaymentMethod(i.PaymentMethod),
		UserID:        i.UserID,
		PhoneNumber:   i.PhoneNumber,
		AddressID:     i.AddressID,
		EmailAddress:  i.EmailAddress,
		ZipCode:       i.ZipCode,
		Prefecture:    i.Prefecture,
		City:          i.City,
		StreetAddress: i.StreetAddress,
		Building:      i.Building,
		LastName:      i.LastName,
		FirstName:     i.FirstName,
	}
}

func getPaymentMethod(m graphql.PaymentMethod) model.PaymentMethod {
	switch m {
	case graphql.PaymentMethodCard:
		return model.PaymentMethodCard
	case graphql.PaymentMethodPaypay:
		return model.PaymentMethodPayPay
	case graphql.PaymentMethodApplePay:
		return model.PaymentMethodApplePay
	case graphql.PaymentMethodGooglePay:
		return model.PaymentMethodGooglePay
	default:
		return model.PaymentMethodUnknown
	}
}
