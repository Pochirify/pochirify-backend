package repository

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

type ProductRepository interface {
	Find(ctx context.Context, id string) (*model.Product, error)
	GetMultiByIDs(ctx context.Context, ids []string) ([]*model.Product, error)
	FindProductVariant(ctx context.Context, id string) (*model.ProductVariant, error)
}
