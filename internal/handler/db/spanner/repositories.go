package spanner

import "github.com/Pochirify/pochirify-backend/internal/domain/repository"

func InitRepositories(spanner *Spanner) repository.Repositories {
	return repository.Repositories{
		UserRepo:    newUserRepository(spanner),
		OrderRepo:   newOrderRepository(spanner),
		ProductRepo: newProductRepository(spanner),
		Tx:          newTransaction(),
	}
}
