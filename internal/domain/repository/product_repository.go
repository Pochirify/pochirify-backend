package repository

type ProductRepository interface {
	// Find(ctx context.Context, id string) (*model.Product, error)
	// GetMultiByIDs(ctx context.Context, ids []string) ([]*model.Product, error)
	// FindProductVariant(ctx context.Context, id string) (*model.ProductVariant, error)
	// Create(ctx context.Context, product *model.Product) error
	// // 現状patchにしてもよさそう
	// Update(ctx context.Context, product *model.Product) error
}
