package shopify

import (
	"bytes"
	"net/http"
)

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
