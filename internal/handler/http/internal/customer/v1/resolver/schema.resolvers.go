package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"errors"
	"fmt"

	graphql1 "github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql/request"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql/response"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/schema"
	"github.com/Pochirify/pochirify-backend/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input graphql1.CreateOrderInput) (*graphql1.CreateOrderPayload, error) {
	output, err := r.App.CreateOrder(
		ctx,
		request.NewCreateOrderInput(input),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	payload, err := response.NewCreateOrderPayload(output)
	if err != nil {
		return nil, fmt.Errorf("failed to create order payload: %w", err)
	}

	return payload, nil
}

// CompleteOrder is the resolver for the completeOrder field.
func (r *mutationResolver) CompleteOrder(ctx context.Context, id string) (*graphql1.CompleteOrderPayload, error) {
	output, err := r.App.CompleteOrder(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrCompleteOrderOrderIsNotCompleted):
			return &graphql1.CompleteOrderPayload{
				ShopifyActivationURL: nil,
				IsNotOrderCompleted:  true,
			}, nil
		default:
			return nil, fmt.Errorf("failed to complete order: %w", err)
		}
	}

	return &graphql1.CompleteOrderPayload{
		ShopifyActivationURL: output.ShopifyActivationURL,
		IsNotOrderCompleted:  false,
	}, nil
}

// TODO: implement
func (r *queryResolver) VariantGroupDetail(ctx context.Context, id string) (*graphql1.VariantGroupDetail, error) {
	return &graphql1.VariantGroupDetail{
		VariantGroup: &graphql1.VariantGroup{
			ID:    id,
			Title: "八天堂のクリームパン",
			ImageURLs: []string{
				"https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
				"https://storage.googleapis.com/pochirify-dev-server-assets/product_images/pic_prod_02%202.png",
			},
			DeliveryTimeRange: &graphql1.DeliveryTimeRange{
				From: "12/11",
				To:   "12/13",
			},
			FaqImageURL:         "https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
			DescriptionImageURL: "https://storage.googleapis.com/pochirify-dev-server-assets/variant_group_descriptions/description.png",
			BadgeImageURL:       "https://storage.googleapis.com/pochirify-dev-server-assets/variang_group_badges/badges.png",
		},
		Variants: []*graphql1.ProductVariant{
			{
				ID:        44794996293943,
				Title:     "お歳暮 ギフトセット",
				UnitPrice: 4800,
				Contents: []string{
					"クリームパン ✖️1",
					"クリームパン 茶色 ✖️2",
				},
				ImageURL: "https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
			},
			{
				ID:        44794997899575,
				Title:     "お歳暮 ギフトセット2",
				UnitPrice: 3800,
				Contents: []string{
					"クリームパン ✖️1",
					"クリームパン 茶色 ✖️2",
				},
				ImageURL: "https://storage.googleapis.com/pochirify-dev-server-assets/product_images/pic_prod_02%202.png",
			},
			{
				ID:        44794998325559,
				Title:     "お歳暮 ギフトセット2",
				UnitPrice: 2800,
				Contents: []string{
					"クリームパン ✖️1",
					"クリームパン 茶色 ✖️2",
				},
				ImageURL: "https://storage.googleapis.com/pochirify-dev-server-assets/product_images/pic_prod_02%202.png",
			},
		},
	}, nil
}

// AllActiveVariantGroupIDs is the resolver for the AllActiveVariantGroupIDs field.
func (r *queryResolver) AllActiveVariantGroupIDs(ctx context.Context) (*graphql1.AllActiveVariantGroupIDs, error) {
	return &graphql1.AllActiveVariantGroupIDs{
		Ids: []string{"1", "2"},
	}, nil
}

// Mutation returns schema.MutationResolver implementation.
func (r *Resolver) Mutation() schema.MutationResolver { return &mutationResolver{r} }

// Query returns schema.QueryResolver implementation.
func (r *Resolver) Query() schema.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
