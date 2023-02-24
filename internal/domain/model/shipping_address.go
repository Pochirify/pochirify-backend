package model

import (
	"time"
)

type ShippingAddress struct {
	// ID string
	// UserID        string
	EmailAddress  EmailAddress
	ZipCode       ZipCode
	Prefecture    string
	City          string
	StreetAddress string
	Building      *string
	LastName      string
	FirstName     string
	PhoneNumber   string
	CreateTime    time.Time
	UpdateTime    time.Time
}

func NewShippingAddress(
	emailAddress string,
	zipCode uint,
	prefecture,
	city,
	streetAddress string,
	building *string,
	lastName,
	firstName,
	phoneNumber string,
) (*ShippingAddress, error) {
	email, err := NewEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}
	zip, err := NewZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	return &ShippingAddress{
		EmailAddress:  email,
		ZipCode:       zip,
		Prefecture:    prefecture,
		City:          city,
		StreetAddress: streetAddress,
		Building:      building,
		LastName:      lastName,
		FirstName:     firstName,
		PhoneNumber:   phoneNumber,
	}, nil
}

type ZipCode uint

// FIXME: properer validation
func NewZipCode(zipCode uint) (ZipCode, error) {
	return ZipCode(zipCode), nil
}

func (z ZipCode) ToUint() uint {
	return uint(z)
}
