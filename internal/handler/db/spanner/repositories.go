package spanner

import "github.com/Pochirify/pochirify-backend/internal/domain/repository"

func InitRepositories() repository.Repositories {
	return repository.Repositories{
		UserRepo:    newUserRepository(),
		OrderRepo:   newOrderRepository(),
		ProductRepo: newProductRepository(),
		Tx:          newTransaction(),
	}
}
