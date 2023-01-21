package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
)

var (
	_ payment.CreditCardClient = (*creditCardClient)(nil)

	endpoint   = "/v1/payments"
	method     = "POST"
	dataFormat = `{
                "pay_type": "Card",
                "job_code": "CAPTURE",
                "amount":  "%d"
            }`
	errCreateCardOrder = errors.New("pochirify-backend-internal-handler-payment-card: failed to create order")
)

type creditCardClient struct {
	ApiKey  string
	BaseURL string
}

func NewCreditCardClient(apiKey, baseURL string) payment.CreditCardClient {
	return &creditCardClient{
		ApiKey:  apiKey,
		BaseURL: baseURL,
	}
}

type orderEntity struct {
	ID       string `json:"id"`
	AccessID string `json:"access_id"`
}

func (e orderEntity) toModel() *payment.CreditCardOrder {
	return &payment.CreditCardOrder{
		CardOrderID: e.ID,
		AccessID:    e.AccessID,
	}
}

func (c creditCardClient) CreateOrder(ctx context.Context, orderID string, amount int) (*payment.CreditCardOrder, error) {
	// https://docs.fincode.jp/api#tag/%E6%B1%BA%E6%B8%88/operation/postPayments
	req, err := http.NewRequest(
		method,
		c.getURL(),
		bytes.NewBuffer(c.getData(amount)),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateCardOrder)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateCardOrder)
	}

	bodyJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errCreateCardOrder)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %w", string(bodyJsonBytes), errCreateCardOrder)
	}

	e := orderEntity{}
	if err = json.Unmarshal(bodyJsonBytes, &e); err != nil {
		return nil, fmt.Errorf("%s: %w", string(bodyJsonBytes), errCreateCardOrder)
	}

	return e.toModel(), nil
}

func (c creditCardClient) getURL() string {
	return c.BaseURL + endpoint
}

func (c creditCardClient) getData(amount int) []byte {
	return []byte(fmt.Sprintf(dataFormat, amount))
}
