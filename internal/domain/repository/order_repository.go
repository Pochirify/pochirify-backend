package repository

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

type OrderRepository interface {
	Find(ctx context.Context, id string) (*model.Order, error)
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
}
