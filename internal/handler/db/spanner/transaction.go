package spanner

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

var _ repository.Transaction = (*transaction)(nil)

type transaction struct {
}

func newTransaction() repository.Transaction {
	return &transaction{}
}

func (t transaction) Transaction(ctx context.Context, f func(context.Context) error) error {
	return nil
}
