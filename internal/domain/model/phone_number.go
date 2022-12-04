package model

type PhoneNumber string

func NewPhoneNumber(phoneNumber string) (PhoneNumber, error) {
	return PhoneNumber(phoneNumber), nil
}

func (ea PhoneNumber) ToDigest() string {
	return generateHashKey(ea.string())
}

func (pn PhoneNumber) string() string {
	return string(pn)
}
