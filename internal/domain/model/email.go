package model

type EmailAddress string

// TODO: validate
func newEmailAddress(address string) (EmailAddress, error) {
	return EmailAddress(address), nil
}

func (ea EmailAddress) toDigest() string {
	return generateHashKey(ea.String())
}

func (ea EmailAddress) String() string {
	return string(ea)
}
