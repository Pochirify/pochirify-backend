package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type queryInfo struct {
	operation string
	name      string
}

func (q *queryInfo) getOperation() string {
	if q == nil {
		return ""
	}
	return q.operation
}

func (q *queryInfo) getName() string {
	if q == nil {
		return ""
	}
	return q.name
}

func getQueryInfo(r *http.Request) (*queryInfo, error) {
	bufOfRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// 消費されてしまったRequest Bodyを修復する
	r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))

	body := string(bufOfRequestBody)
	splitted := strings.Split(body, " ")
	if len(splitted) < 2 {
		return nil, fmt.Errorf("unexpected body: %s", body)
	}

	return &queryInfo{
		operation: splitted[0],
		name:      splitted[1],
	}, nil
}
