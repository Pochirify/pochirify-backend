package model

// var newUserError = errors.New("pochirify-backend-internal-domain-model: failed to create new user")

// type User struct {
// 	// TODO: userIDはphoneNumberからつくる。住所は複数持つ
// 	ID                string
// 	PhoneNumberDigest string
// 	IsAuthenticated   bool
// 	CreateTime        time.Time
// 	UpdateTime        time.Time
// }

// func NewUser(phoneNumber string) (*User, error) {
// 	pn, err := newPhoneNumber(phoneNumber)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, newUserError)
// 	}

// 	return &User{
// 		ID:                uuid.NewString(),
// 		PhoneNumberDigest: pn.toDigest(),
// 		IsAuthenticated:   false,
// 	}, nil
// }
