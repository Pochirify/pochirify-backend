package model

type EmailAddress string

func NewEmailAddress(address string) (EmailAddress, error) {
	return EmailAddress(address), nil
}

func (ea EmailAddress) ToDigest() string {
	return generateHashKey(ea.string())
}

func (ea EmailAddress) string() string {
	return string(ea)
}
