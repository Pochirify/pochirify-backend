package spanner

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner/yo"
	"github.com/google/uuid"
)

var _ repository.ProductRepository = (*productRepository)(nil)

type productRepository struct {
	*Spanner
}

func newProductRepository(s *Spanner) repository.ProductRepository {
	return &productRepository{s}
}

type productEntity yo.Product

func newProductEntity(m *model.Product) *productEntity {
	now := time.Now()
	if m.CreateTime.IsZero() {
		m.CreateTime = now
	}
	m.UpdateTime = now

	return &productEntity{
		ID:               m.ID,
		Title:            m.Title,
		Price:            int64(m.Price),
		Stock:            int64(m.Stock),
		ContentOne:       m.ContentOne,
		ContentTwo:       toSpannerNullString(m.ContentTwo),
		ContentThree:     toSpannerNullString(m.ContentThree),
		ContentFour:      toSpannerNullString(m.ContentFour),
		ContentFive:      toSpannerNullString(m.ContentFive),
		ProductVariantID: toSpannerNullString(m.ProductVariantID),
		CreateTime:       m.CreateTime,
		UpdateTime:       m.UpdateTime,
	}
}

func (e *productEntity) toModel() (*model.Product, error) {
	return &model.Product{
		ID:               e.ID,
		Title:            e.Title,
		Price:            int(e.Price),
		Stock:            int(e.Stock),
		ContentOne:       e.ContentOne,
		ContentTwo:       fromSpannerNullString(e.ContentTwo),
		ContentThree:     fromSpannerNullString(e.ContentThree),
		ContentFour:      fromSpannerNullString(e.ContentFour),
		ContentFive:      fromSpannerNullString(e.ContentFive),
		ProductVariantID: fromSpannerNullString(e.ProductVariantID),
		CreateTime:       e.CreateTime,
		UpdateTime:       e.UpdateTime,
	}, nil
}

func (r *productRepository) Find(ctx context.Context, id string) (*model.Product, error) {
	yo, err := yo.FindProduct(ctx, r.Ctx(ctx), id)
	if err != nil {
		switch {
		case isNotFoundErr(err):
			return nil, findError([]field{{"productID", id}}, err, model.NotFoundError)
		default:
			return nil, findError([]field{{"productID", id}}, err)
		}
	}

	return (*productEntity)(yo).toModel()
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	e := newProductEntity(product)
	mutation := (*yo.Product)(e).Update(ctx)
	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetMultiByIDs(ctx context.Context, ids []string) ([]*model.Product, error) {
	return []*model.Product{
		{
			ID:    "1",
			Title: "商品タイトルるううううううううううううう",
			Price: 100,
		},
		{
			ID:    "2",
			Title: "商品タイトル2222222222222222",
			Price: 100,
		},
	}, nil
}

func (r *productRepository) FindProductVariant(ctx context.Context, id string) (*model.ProductVariant, error) {
	return &model.ProductVariant{
		ID:         uuid.NewString(),
		ProductIDs: [5]string{"1", "2"},
	}, nil
}
