package model

import "time"

type Product struct {
	ID string
	// VendorID         string
	Title            string
	Price            int
	Stock            int
	ContentOne       string
	ContentTwo       *string
	ContentThree     *string
	ContentFour      *string
	ContentFive      *string
	ProductVariantID *string
	CreateTime       time.Time
	UpdateTime       time.Time
}

func (p *Product) Bought() {
	p.Stock--
}

func (p Product) GetTotalPrice(quantity int) int {
	return p.Price * quantity
}
