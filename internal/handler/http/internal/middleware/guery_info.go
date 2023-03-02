package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type queryInfo struct {
	operation     string
	operationName string
}

func (q *queryInfo) getOperation() string {
	if q == nil {
		return ""
	}
	return q.operation
}

func (q *queryInfo) getOperationName() string {
	if q == nil {
		return ""
	}
	return q.operationName
}

type requestBody struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
}

func getQueryInfo(r *http.Request) (*queryInfo, error) {
	bufOfRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// 消費されてしまったRequest Bodyを修復する
	r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))

	var requestBody requestBody
	if err = json.Unmarshal(bufOfRequestBody, &requestBody); err != nil {
		return nil, err
	}

	splitted := strings.Split(requestBody.Query, " ")
	if len(splitted) < 2 {
		return nil, errors.New("unexpected body")
	}

	return &queryInfo{
		operation:     splitted[0],
		operationName: requestBody.OperationName,
	}, nil
}
