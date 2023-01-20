package spanner

import (
	"context"
	"log"

	"github.com/Pochirify/pochirify-backend/internal/domain/model"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

var _ repository.UserRepository = (*userRepository)(nil)

type userRepository struct{}

func newUserRepository() *userRepository {
	return &userRepository{}
}

func (r userRepository) Find(ctx context.Context, userID string) (*model.User, error) {
	return nil, nil
}

func (r userRepository) Create(ctx context.Context, u *model.User) error {
	return nil
}

func (r userRepository) FindUserAddress(ctx context.Context, addressID string) (*model.UserAddress, error) {
	return nil, nil
}

func (r userRepository) Upsert(ctx context.Context, user *model.User) error {
	log.Println(user)
	return nil
}

func (r userRepository) CreateUserAddress(ctx context.Context, userAddress *model.UserAddress) error {
	log.Println(userAddress)
	return nil
}
