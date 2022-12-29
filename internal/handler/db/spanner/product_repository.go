package spanner

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
	"github.com/google/uuid"
)

var _ repository.ProductRepository = (*productRepository)(nil)

type productRepository struct {
}

func newProductRepository() repository.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Find(ctx context.Context, id string) (*model.Product, error) {
	vID := "variantID"
	return &model.Product{
		ID:               id,
		Title:            "商品タイトルるううううううううううううう",
		Price:            100,
		Contents:         [5]string{"contents1", "contents2", "contents3", "contents4", "contents5"},
		ProductVariantID: &vID,
	}, nil
}

func (r *productRepository) GetMultiByIDs(ctx context.Context, ids []string) ([]*model.Product, error) {
	return []*model.Product{
		{
			ID:       "1",
			Title:    "商品タイトルるううううううううううううう",
			Price:    100,
			Contents: [5]string{"contents1", "contents2", "contents3", "contents4"},
		},
		{
			ID:       "2",
			Title:    "商品タイトル2222222222222222",
			Price:    100,
			Contents: [5]string{"contents1", "contents2", "contents3"},
		},
	}, nil
}

func (r *productRepository) FindProductVariant(ctx context.Context, id string) (*model.ProductVariant, error) {
	return &model.ProductVariant{
		ID:         uuid.NewString(),
		ProductIDs: [5]string{"1", "2"},
	}, nil
}
