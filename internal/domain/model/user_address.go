package model

// // User : UserAddress = 1 : N
// type UserAddress struct {
// 	ID            string
// 	UserID        string
// 	EmailAddress  EmailAddress
// 	ZipCode       ZipCode
// 	Prefecture    string
// 	City          string
// 	StreetAddress string
// 	Building      *string
// 	LastName      string
// 	FirstName     string
// 	CreateTime    time.Time
// 	UpdateTime    time.Time
// }

// func NewUserAddress(
// 	userID,
// 	emailAddress string,
// 	zipCode int,
// 	prefecture,
// 	city,
// 	streetAddress string,
// 	building *string,
// 	lastName,
// 	firstName string,
// ) (*UserAddress, error) {
// 	email, err := NewEmailAddress(emailAddress)
// 	if err != nil {
// 		return nil, err
// 	}
// 	zip, err := NewZipCode(zipCode)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &UserAddress{
// 		ID:            uuid.NewString(),
// 		UserID:        userID,
// 		EmailAddress:  email,
// 		ZipCode:       zip,
// 		Prefecture:    prefecture,
// 		City:          city,
// 		StreetAddress: streetAddress,
// 		Building:      building,
// 		LastName:      lastName,
// 		FirstName:     firstName,
// 	}, nil
// }

// type ZipCode int

// // FIXME: properer validation
// func NewZipCode(zipCode int) (ZipCode, error) {
// 	return ZipCode(zipCode), nil
// }

// func (z ZipCode) ToInt() int {
// 	return int(z)
// }
