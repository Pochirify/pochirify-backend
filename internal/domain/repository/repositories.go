package repository

type Repositories struct {
	UserRepo    UserRepository
	OrderRepo   OrderRepository
	ProductRepo ProductRepository
	Tx          Transaction
}
