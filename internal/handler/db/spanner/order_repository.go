package spanner

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner/yo"
)

var _ repository.OrderRepository = (*orderRepository)(nil)

type orderRepository struct {
	*Spanner
}

func newOrderRepository(s *Spanner) repository.OrderRepository {
	return &orderRepository{s}
}

type orderEntity yo.Order

func newOrderEntity(m *model.Order) *orderEntity {
	now := time.Now()
	if m.CreateTime.IsZero() {
		m.CreateTime = now
	}
	m.UpdateTime = now

	return &orderEntity{
		ID:               m.ID,
		ShopifyOrderID:   int64(m.ShopifyOrderID),
		Status:           m.Status.String(),
		PaymentMethod:    m.PaymentMethod.String(),
		ProductVariantID: int64(m.ProductVariantID),
		UnitPrice:        int64(m.UnitPrice),
		Quantity:         int64(m.Quantity),
		CreateTime:       m.CreateTime,
		UpdateTime:       m.UpdateTime,
	}
}

func (e *orderEntity) toModel() (*model.Order, error) {
	return &model.Order{
		e.ID,
		uint(e.ShopifyOrderID),
		model.GetOrderPaymentStatus(e.Status),
		model.GetPaymentMethod(e.PaymentMethod),
		uint(e.ProductVariantID),
		uint(e.UnitPrice),
		uint(e.Quantity),
		e.CreateTime,
		e.UpdateTime,
	}, nil
}

func (r *orderRepository) Find(ctx context.Context, id string) (*model.Order, error) {
	yo, err := yo.FindOrder(ctx, r.Ctx(ctx), id)
	if err != nil {
		switch {
		case isNotFoundErr(err):
			return nil, findError([]field{{"orderID", id}}, err, model.NotFoundError)
		default:
			return nil, findError([]field{{"orderID", id}}, err)
		}
	}

	return (*orderEntity)(yo).toModel()
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	e := newOrderEntity(order)
	mutation := (*yo.Order)(e).Insert(ctx)
	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Update(ctx context.Context, order *model.Order) error {
	e := newOrderEntity(order)
	mutation := (*yo.Order)(e).Update(ctx)
	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
		return err
	}

	return nil
}
