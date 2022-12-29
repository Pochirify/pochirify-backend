package model

import (
	"time"
)

type ProductVariantMaster struct {
	ID string
}

type ProductVariant struct {
	ID         string
	ProductIDs [5]string
	// ProductIDOne   string
	// ProductIDTwo   string
	// ProductIDThree *string
	// ProductIDFour  *string
	// ProductIDFive  *string
	CreateTime time.Time
	UpdateTime time.Time
}

func (pv *ProductVariant) GetProductIDsExcept(exceptID string) []string {
	productIDs := make([]string, 0, 5)
	if pv.ProductIDs[0] != "" && pv.ProductIDs[0] != exceptID {
		productIDs = append(productIDs, pv.ProductIDs[0])
	}
	if pv.ProductIDs[1] != "" && pv.ProductIDs[1] != exceptID {
		productIDs = append(productIDs, pv.ProductIDs[1])
	}
	if pv.ProductIDs[2] != "" && pv.ProductIDs[2] != exceptID {
		productIDs = append(productIDs, pv.ProductIDs[0])
	}
	if pv.ProductIDs[3] != "" && pv.ProductIDs[3] != exceptID {
		productIDs = append(productIDs, pv.ProductIDs[3])
	}
	if pv.ProductIDs[4] != "" && pv.ProductIDs[4] != exceptID {
		productIDs = append(productIDs, pv.ProductIDs[4])
	}

	return productIDs
}
