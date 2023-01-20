package model

type PhoneNumber string

// TODO: validate
func newPhoneNumber(phoneNumber string) (PhoneNumber, error) {
	return PhoneNumber(phoneNumber), nil
}

func (ea PhoneNumber) toDigest() string {
	return generateHashKey(ea.String())
}

func (pn PhoneNumber) String() string {
	return string(pn)
}
