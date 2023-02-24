package shopify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

var errorMessageUnableToReserveInventory = "Unable to reserve inventory"

type LineItem struct {
	Quantity  uint   `json:"quantity"`
	VariantID string `json:"variant_id"`
}
type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type BillingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address2  string `json:"address2"`
	Address1  string `json:"address1"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Zip       int    `json:"zip"`
	Phone     string `json:"phone"`
}
type ShippingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address2  string `json:"address2"`
	Address1  string `json:"address1"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Zip       int    `json:"zip"`
	Phone     string `json:"phone"`
}

type ReqOrder struct {
	InventoryBehaviour string          `json:"inventory_behaviour"`
	FinancialStatus    string          `json:"financial_status"`
	LineItems          []LineItem      `json:"line_items"`
	Customer           Customer        `json:"customer"`
	BillingAddress     BillingAddress  `json:"billing_address"`
	ShippingAddress    ShippingAddress `json:"shipping_address"`
}

type createOrderRequest struct {
	Order ReqOrder `json:"order"`
}

type ResOrder struct {
	Order struct {
		ID         int    `json:"id"`
		TotalPrice string `json:"total_price"`
	} `json:"order"`
}

type CreateOrderError struct {
	Errors struct {
		LineItems []string `json:"line_items"`
	} `json:"errors"`
}

func (o ResOrder) toPayload(order *model.Order) (*shopify.CreateOrderPayload, error) {
	expectedTotalPrice := order.Quantity * order.UnitPrice
	totalPrice, err := strconv.Atoi(o.Order.TotalPrice)
	if err != nil {
		return nil, err
	}
	isPriceChanged := expectedTotalPrice != uint(totalPrice)
	return &shopify.CreateOrderPayload{
		ShopifyOrderID: o.Order.ID,
		IsPriceChanged: isPriceChanged,
		TotalPrice:     uint(totalPrice),
	}, nil
}

func (client shopifyClient) CreateOrder(
	ctx context.Context,
	order *model.Order,
	shippingAddress *model.ShippingAddress,
) (*shopify.CreateOrderPayload, error) {
	var address2 string
	if shippingAddress.Building != nil {
		address2 = *shippingAddress.Building
	}
	body := &createOrderRequest{
		Order: ReqOrder{
			InventoryBehaviour: "decrement_obeying_policy",
			FinancialStatus:    "pending",
			LineItems: []LineItem{
				{
					Quantity:  order.Quantity,
					VariantID: order.ProductVariantID,
				},
			},
			Customer: Customer{
				FirstName: shippingAddress.FirstName,
				LastName:  shippingAddress.LastName,
				Email:     shippingAddress.EmailAddress.String(),
			},
			// TODO: 同じ変数に詰めてからわたした方がbetter
			BillingAddress: BillingAddress{
				FirstName: shippingAddress.FirstName,
				LastName:  shippingAddress.LastName,
				Address2:  address2,
				Address1:  shippingAddress.StreetAddress,
				City:      shippingAddress.City,
				Province:  shippingAddress.Prefecture,
				Country:   "Japan",
				Zip:       int(shippingAddress.ZipCode.ToUint()),
				Phone:     shippingAddress.PhoneNumber,
			},
			ShippingAddress: ShippingAddress{
				FirstName: shippingAddress.FirstName,
				LastName:  shippingAddress.LastName,
				Address2:  address2,
				Address1:  shippingAddress.StreetAddress,
				City:      shippingAddress.City,
				Province:  shippingAddress.Prefecture,
				Country:   "Japan",
				Zip:       int(shippingAddress.ZipCode.ToUint()),
				Phone:     shippingAddress.PhoneNumber,
			},
		},
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := client.doPostRequest("orders.json", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		var resOrder ResOrder
		if err = json.Unmarshal(resBody, &resOrder); err != nil {
			return nil, err
		}
		return resOrder.toPayload(order)
	case http.StatusUnprocessableEntity:
		var createOrderError CreateOrderError
		if err = json.Unmarshal(resBody, &createOrderError); err != nil {
			return nil, err
		}
		if len(createOrderError.Errors.LineItems) <= 0 {
			break
		}

		if strings.Contains(createOrderError.Errors.LineItems[0], errorMessageUnableToReserveInventory) {
			return nil, shopify.ErrCreateOrderUnableToReserveInventory
		}
	}

	// TODO: logにrespのメッセージを出すべき
	return nil, fmt.Errorf("unexpected create order error occurred. statusCode=%d", resp.StatusCode)
}
