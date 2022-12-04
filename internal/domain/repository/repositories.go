package repository

type Repositories struct {
	UserRepo  UserRepository
	OrderRepo OrderRepository
	Tx        Transaction
}
