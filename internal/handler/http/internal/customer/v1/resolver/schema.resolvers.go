package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graqhql"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/schema"
	"github.com/Pochirify/pochirify-backend/internal/usecase"
	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input graqhql.NewTodo) (*graqhql.Todo, error) {
	todo := graqhql.Todo{
		ID:   uuid.NewString(),
		Text: input.Text,
		Done: false,
	}
	return &todo, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input graqhql.NewUser) (*graqhql.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreatePaypayQRCode is the resolver for the createPaypayQRCode field.
func (r *mutationResolver) CreatePaypayQRCode(ctx context.Context, input graqhql.PaypayQRCodeInput) (*graqhql.CreatePaypayQRCodePayload, error) {
	output, err := r.App.CreatePaypayQRCode(ctx, &usecase.CreatePaypayQRCodeInput{
		EmailAddress: input.EmailAddress,
		PhoneNumber:  input.PhoneNumber,
		Zip:          input.Address.Zip,
		Prefecture:   input.Address.Prefecture,
		AddressOne:   input.Address.AddressOne,
		AddressTwo:   input.Address.AddressTwo,

		Amount:           input.Amount,
		OrderDescription: input.OrderDescription,
	})
	if err != nil {
		return nil, err
	}

	return &graqhql.CreatePaypayQRCodePayload{
		URL:      output.QRCode.QRCodeUrl,
		DeepLink: output.QRCode.QRCodeDeepLink,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*graqhql.Todo, error) {
	return []*graqhql.Todo{
		{Text: "text"},
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*graqhql.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *graqhql.Todo) (*graqhql.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Todo is the resolver for the todo field.
func (r *userResolver) Todo(ctx context.Context, obj *graqhql.User) (*graqhql.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// Mutation returns schema.MutationResolver implementation.
func (r *Resolver) Mutation() schema.MutationResolver { return &mutationResolver{r} }

// Query returns schema.QueryResolver implementation.
func (r *Resolver) Query() schema.QueryResolver { return &queryResolver{r} }

// Todo returns schema.TodoResolver implementation.
func (r *Resolver) Todo() schema.TodoResolver { return &todoResolver{r} }

// User returns schema.UserResolver implementation.
func (r *Resolver) User() schema.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
