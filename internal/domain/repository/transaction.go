package repository

import "context"

type Transaction interface {
	Transaction(ctx context.Context, f func(context.Context) error) error
}
