package model

import "time"

type Product struct {
	ID                string
	VendorID          string
	Title             string
	Price             uint
	Contents          [5]string
	ShippingTimeRange *ShippingTimeRange
	ProductVariantID  *string
	CreateTime        time.Time
	UpdateTime        time.Time
}
