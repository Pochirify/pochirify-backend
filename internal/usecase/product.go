package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

var (
	errGetProductDetail = errors.New("")
)

type GetProductDetailInput struct {
	ProductID string
}

type GetProductDetailOutput struct {
	Product  *model.Product
	Variants []*model.Product
}

func (a App) GetProductDetail(ctx context.Context, input *GetProductDetailInput) (*GetProductDetailOutput, error) {
	p, err := a.ProductRepo.Find(ctx, input.ProductID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errGetProductDetail)
	}
	if p.ProductVariantID == nil {
		return &GetProductDetailOutput{
			Product:  p,
			Variants: nil,
		}, nil
	}

	pv, err := a.ProductRepo.FindProductVariant(ctx, p.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errGetProductDetail)
	}
	productIDs := pv.GetProductIDsExcept(p.ID)
	products, err := a.ProductRepo.GetMultiByIDs(ctx, productIDs)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, errGetProductDetail)
	}

	return &GetProductDetailOutput{
		Product:  p,
		Variants: products,
	}, nil
}
