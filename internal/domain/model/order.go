package model

import "github.com/google/uuid"

type Order struct {
	ID               string
	UserID           UserID
	UserAddressID    string
	Amount           int
	OrderDescription string
}

func NewOrder(user *User, userAddress *UserAddress, amount int, orderDescription string) *Order {
	return &Order{
		ID:               uuid.NewString(),
		UserID:           user.ID,
		UserAddressID:    userAddress.ID,
		Amount:           amount,
		OrderDescription: orderDescription,
	}
}
