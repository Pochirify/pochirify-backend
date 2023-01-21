package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	graphql1 "github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/graphql/request"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/schema"
)

// TODO: error文はinternalで閉じる
// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input graphql1.CreateOrderInput) (*graphql1.CreateOrderPayload, error) {
	output, err := r.App.CreateOrder(
		ctx,
		request.NewCreateOrderInput(input),
	)
	if err != nil {
		return nil, err
	}

	return &graphql1.CreateOrderPayload{
		OrderID: output.OrderID,
		URL:     output.URL,
	}, nil
}

// VariantGroupDetail is the resolver for the variantGroupDetail field.
func (r *queryResolver) VariantGroupDetail(ctx context.Context, id string) (*graphql1.VariantGroupDetail, error) {
	return &graphql1.VariantGroupDetail{
		VariantGroup: &graphql1.VariantGroup{
			ID:    id,
			Title: "八天堂の商品たち",
			ImageURLs: []string{
				"https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
				"https://storage.googleapis.com/pochirify-dev-server-assets/product_images/pic_prod_02%202.png",
			},
			DeliveryTimeRange: &graphql1.DeliveryTimeRange{
				From: "12/11",
				To:   "12/13",
			},
			FaqImageURL: &graphql1.WebpPngImageURL{
				WebpURL: "https://storage.googleapis.com/adfsafdafd/FAQ.webp",
				PngURL:  "https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
			},
			DescriptionImageURL: &graphql1.WebpPngImageURL{
				WebpURL: "https://storage.googleapis.com/pochirify-dev-server-assets/variant_group_descriptions/description.webp",
				PngURL:  "https://storage.googleapis.com/pochirify-dev-server-assets/variant_group_descriptions/description.png",
			},
			BadgeImageURL: "https://storage.googleapis.com/pochirify-dev-server-assets/variang_group_badges/badges.png",
		},
		Variants: []*graphql1.Product{
			{
				ID:    "1",
				Title: "お歳暮 ギフトセット",
				Price: 4800,
				Contents: []string{
					"クリームパン ✖️1",
					"クリームパン 茶色 ✖️2",
				},
				ImageURL: "https://storage.googleapis.com/adfsafdafd/image%2010%20(1).png",
			},
			{
				ID:    "2",
				Title: "お歳暮 ギフトセット2",
				Price: 3800,
				Contents: []string{
					"クリームパン ✖️1",
					"クリームパン 茶色 ✖️2",
				},
				ImageURL: "https://storage.googleapis.com/pochirify-dev-server-assets/product_images/pic_prod_02%202.png",
			},
			{
				ID:    "3",
				Title: "お歳暮 ギフトセット2",
				Price: 2800,
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
