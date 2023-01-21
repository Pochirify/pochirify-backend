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
		ID:            m.ID,
		UserID:        m.UserID,
		UserAddressID: m.UserAddressID,
		Status:        m.Status.String(),
		PaymentMethod: m.PaymentMethod.String(),
		ProductID:     m.ProductID,
		Price:         int64(m.Price),
		CreateTime:    m.CreateTime,
		UpdateTime:    m.UpdateTime,
	}
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	e := newOrderEntity(order)
	mutation := (*yo.Order)(e).Insert(ctx)
	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
		return err
	}

	return nil
}
