package response

// func NewGetProductDetailRes(o *usecase.GetProductDetailOutput) *graphql.ProductDetail {
// 	p := newProduct(o.Product)
// 	variants := make([]*graphql.Product, 0, len(o.Variants))
// 	for _, v := range o.Variants {
// 		variants = append(variants, newProduct(v))
// 	}
// 	return &graphql.ProductDetail{
// 		Product:  p,
// 		Variants: variants,
// 	}
// }

// func newProduct(product *model.Product) *graphql.Product {
// 	contents := make([]string, 0, len(product.Contents))
// 	for _, content := range product.Contents {
// 		if content == "" {
// 			break
// 		}
// 		contents = append(contents, content)
// 	}
// 	return &graphql.Product{
// 		ID:       product.ID,
// 		Title:    product.Title,
// 		Price:    int(product.Price),
// 		Contents: contents,
// 		RoughTimeRange: &graphql.RoughTimeRange{
// 			From: "12/10",
// 			To:   "12/13",
// 		},
// 	}
// }
