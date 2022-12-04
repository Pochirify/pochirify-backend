package spanner

import (
	"context"
	"log"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

var _ repository.OrderRepository = (*orderRepository)(nil)

type orderRepository struct {
}

func newOrderRepository() repository.OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	log.Println(order)
	return nil
}
