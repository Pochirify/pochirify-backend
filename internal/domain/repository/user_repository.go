package repository

import (
	"context"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
)

type UserRepository interface {
	Upsert(ctx context.Context, user *model.User) error
	CreateUserAddress(ctx context.Context, userAddress *model.UserAddress) error
}
