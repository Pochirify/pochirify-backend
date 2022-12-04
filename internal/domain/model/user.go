package model

import (
	"errors"
	"fmt"
	"time"
)

var newUserError = errors.New("pochirify-backend-internal-domain-model: failed to create new user")

type User struct {
	ID                 UserID
	EmailAddressDigest string
	PhoneNumberDigest  string
	CreateTime         time.Time
	UpdateTime         time.Time
}

type UserID string

func NewUser(emailAddress, phoneNumber string) (*User, error) {
	ea, err := NewEmailAddress(emailAddress)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, newUserError)
	}
	pn, err := NewPhoneNumber(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, newUserError)
	}

	return &User{
		ID:                 NewUserID(ea, pn),
		EmailAddressDigest: ea.ToDigest(),
		PhoneNumberDigest:  pn.ToDigest(),
	}, nil
}

func NewUserID(emailAddress EmailAddress, phoneNumber PhoneNumber) UserID {
	return UserID(generateHashKey(emailAddress.string(), phoneNumber.string()))
}

func (uid UserID) String() string {
	return string(uid)
}
