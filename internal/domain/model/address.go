package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// User : UserAddress = 1 : N
type UserAddress struct {
	ID         string
	UserID     UserID
	Zip        uint
	Prefecture string
	AddressOne string
	AddressTwo *string
	CreateTime time.Time
	UpdateTime time.Time
}

func NewUserAddress(userID UserID, zip, prefecture, addressOne string, addressTwo *string) (*UserAddress, error) {
	uiZip, err := newZip(zip)
	if err != nil {
		return nil, err
	}

	return &UserAddress{
		ID:         uuid.NewString(),
		UserID:     userID,
		Zip:        uiZip,
		Prefecture: prefecture,
		AddressOne: addressOne,
		AddressTwo: addressTwo,
	}, nil
}

// FIXME: properer validation
func newZip(zip string) (uint, error) {
	splitted := strings.Split(zip, "-")
	switch len(splitted) {
	case 1:
		uiZip, err := strconv.ParseUint(splitted[0], 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uiZip), nil
	case 2:
		uiZip1, err := strconv.ParseUint(splitted[0], 10, 32)
		if err != nil {
			return 0, err
		}
		uiZip2, err := strconv.ParseUint(splitted[0], 10, 32)
		if err != nil {
			return 0, err
		}

		return uint(10000*uiZip1 + uiZip2), nil
	default:
		return 0, fmt.Errorf("unexpected zip. %s", zip)
	}
}
