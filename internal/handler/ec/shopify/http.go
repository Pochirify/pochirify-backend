package shopify

import (
	"bytes"
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

var (
	errorMessageUnableToReserveInventory = "Unable to reserve inventory"
	errorAccountAlreadyEnabled           = "account already enabled"
	postActivateAccountPathPattern       = "customers/%d/account_activation_url.json"
)

type LineItem struct {
	Quantity  uint `json:"quantity"`
	VariantID uint `json:"variant_id"`
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

type createPendingOrderRequest struct {
	Order ReqOrder `json:"order"`
}

type ResOrder struct {
	Order struct {
		ID         int    `json:"id"`
		TotalPrice string `json:"total_price"`
	} `json:"order"`
}

type CreatePendingOrderError struct {
	Errors struct {
		LineItems []string `json:"line_items"`
	} `json:"errors"`
}

func (o ResOrder) toPayload() (*shopify.CreatePendingOrderPayload, error) {
	totalPrice, err := strconv.Atoi(o.Order.TotalPrice)
	if err != nil {
		return nil, err
	}
	return &shopify.CreatePendingOrderPayload{
		ShopifyOrderID: o.Order.ID,
		TotalPrice:     uint(totalPrice),
	}, nil
}

func (client shopifyClient) CreatePendingOrder(
	ctx context.Context,
	quantity,
	productVariantID uint,
	shippingAddress *model.ShippingAddress,
) (*shopify.CreatePendingOrderPayload, error) {
	var address2 string
	if shippingAddress.Building != nil {
		address2 = *shippingAddress.Building
	}
	body := &createPendingOrderRequest{
		Order: ReqOrder{
			InventoryBehaviour: "decrement_obeying_policy",
			FinancialStatus:    "pending",
			LineItems: []LineItem{
				{
					Quantity:  quantity,
					VariantID: productVariantID,
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
		return resOrder.toPayload()
	case http.StatusUnprocessableEntity:
		var createPendingOrderError CreatePendingOrderError
		if err = json.Unmarshal(resBody, &createPendingOrderError); err != nil {
			return nil, err
		}
		if len(createPendingOrderError.Errors.LineItems) == 0 {
			break
		}

		if strings.Contains(createPendingOrderError.Errors.LineItems[0], errorMessageUnableToReserveInventory) {
			return nil, shopify.ErrCreatePendingOrderUnableToReserveInventory
		}
	}

	client.logger(ctx).Error(fmt.Errorf(string(resBody)), "failed to create order")
	return nil, fmt.Errorf("unexpected create order error occurred. statusCode=%d", resp.StatusCode)
}

type ActivationURL struct {
	AccountActivationURL string `json:"account_activation_url"`
}

type GetShopifyActivationURLError struct {
	Errors []string `json:"errors"`
}

func (c shopifyClient) GetShopifyActivationURL(ctx context.Context, shopifyCustomerID uint) (string, error) {
	path := fmt.Sprintf(postActivateAccountPathPattern, shopifyCustomerID)
	resp, err := c.doPostRequest(path, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var activationURL ActivationURL
		if err = json.Unmarshal(resBody, &activationURL); err != nil {
			return "", err
		}
		return activationURL.AccountActivationURL, nil
	case http.StatusUnprocessableEntity:
		var getShopifyActivationURLError GetShopifyActivationURLError
		if err = json.Unmarshal(resBody, &getShopifyActivationURLError); err != nil {
			return "", err
		}
		if len(getShopifyActivationURLError.Errors) == 0 {
			break
		}

		if strings.Contains(getShopifyActivationURLError.Errors[0], errorAccountAlreadyEnabled) {
			return "", shopify.ErrGetShopifyActivationURLAlreadyActivated
		}
	}
	// TODO: logにrespのメッセージを出すべき
	return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}

func (c shopifyClient) doPostRequest(path string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.restBaseURL+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token", c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
