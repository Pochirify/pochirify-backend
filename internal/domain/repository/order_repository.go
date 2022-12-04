package repository

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
}
