package model

import "time"

type Product struct {
	ID                string
	VendorID          string
	Title             string
	Price             int
	ShippingTimeRange int // day
	Stock             int
	ContentOne        string
	ContentTwo        *string
	ContentThree      *string
	ContentFour       *string
	ProductVariantID  *string
	CreateTime        time.Time
	UpdateTime        time.Time
}

func (p *Product) Bought() {
	p.Stock--
}
