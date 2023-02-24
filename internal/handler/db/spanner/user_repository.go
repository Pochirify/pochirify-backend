package spanner

// var (
// 	_ repository.UserRepository = (*userRepository)(nil)

// 	errToUserAddressModel = errors.New("pochirify-backend-internal-handler-db-spanner-user: failed to convert to userAddress model")
// )

// type userRepository struct {
// 	*Spanner
// }

// func newUserRepository(s *Spanner) *userRepository {
// 	return &userRepository{s}
// }

// type userEntity yo.User

// func newUserEntity(m *model.User) *userEntity {
// 	now := time.Now()
// 	if m.CreateTime.IsZero() {
// 		m.CreateTime = now
// 	}
// 	m.UpdateTime = now

// 	return &userEntity{
// 		ID:                m.ID,
// 		PhoneNumberDigest: m.PhoneNumberDigest,
// 		IsAuthenticated:   m.IsAuthenticated,
// 		CreateTime:        m.CreateTime,
// 		UpdateTime:        m.UpdateTime,
// 	}
// }

// func (e *userEntity) toModel() (*model.User, error) {
// 	return &model.User{
// 		ID:                e.ID,
// 		PhoneNumberDigest: e.PhoneNumberDigest,
// 		IsAuthenticated:   e.IsAuthenticated,
// 		CreateTime:        e.CreateTime,
// 		UpdateTime:        e.UpdateTime,
// 	}, nil
// }

// func (r userRepository) Find(ctx context.Context, userID string) (*model.User, error) {
// 	yo, err := yo.FindUser(ctx, r.Ctx(ctx), userID)
// 	if err != nil {
// 		switch {
// 		case isNotFoundErr(err):
// 			return nil, findError([]field{{"userID", userID}}, err, model.NotFoundError)
// 		default:
// 			return nil, findError([]field{{"userID", userID}}, err)
// 		}
// 	}

// 	return (*userEntity)(yo).toModel()
// }

// func (r userRepository) Create(ctx context.Context, u *model.User) error {
// 	e := newUserEntity(u)
// 	mutation := (*yo.User)(e).Insert(ctx)
// 	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UserAddress
// type userAddressEntity yo.UserAddress

// func newUserAddressEntity(m *model.UserAddress) *userAddressEntity {
// 	now := time.Now()
// 	if m.CreateTime.IsZero() {
// 		m.CreateTime = now
// 	}
// 	m.UpdateTime = now

// 	return &userAddressEntity{
// 		ID:            m.ID,
// 		UserID:        m.UserID,
// 		EmailAddress:  m.EmailAddress.String(),
// 		ZipCode:       int64(m.ZipCode.ToInt()),
// 		Prefecture:    m.Prefecture,
// 		City:          m.City,
// 		StreetAddress: m.StreetAddress,
// 		Building:      toSpannerNullString(m.Building),
// 		LastName:      m.LastName,
// 		FirstName:     m.FirstName,
// 		CreateTime:    m.CreateTime,
// 		UpdateTime:    m.UpdateTime,
// 	}
// }

// func (e *userAddressEntity) toModel() (*model.UserAddress, error) {
// 	emailAddress, err := model.NewEmailAddress(e.EmailAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errToUserAddressModel)
// 	}
// 	zipCode, err := model.NewZipCode(int(e.ZipCode))
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", err, errToUserAddressModel)
// 	}
// 	return &model.UserAddress{
// 		ID:            e.ID,
// 		UserID:        e.UserID,
// 		EmailAddress:  emailAddress,
// 		ZipCode:       zipCode,
// 		Prefecture:    e.Prefecture,
// 		City:          e.City,
// 		StreetAddress: e.StreetAddress,
// 		Building:      fromSpannerNullString(e.Building),
// 		LastName:      e.LastName,
// 		FirstName:     e.FirstName,
// 		CreateTime:    e.CreateTime,
// 		UpdateTime:    e.UpdateTime,
// 	}, nil
// }

// func (r userRepository) FindUserAddress(ctx context.Context, addressID string) (*model.UserAddress, error) {
// 	yo, err := yo.FindUserAddress(ctx, r.Ctx(ctx), addressID)
// 	if err != nil {
// 		switch {
// 		case isNotFoundErr(err):
// 			return nil, findError([]field{{"userAddressID", addressID}}, err, model.NotFoundError)
// 		default:
// 			return nil, findError([]field{{"userAddressID", addressID}}, err)
// 		}
// 	}

// 	return (*userAddressEntity)(yo).toModel()
// }

// func (r userRepository) Upsert(ctx context.Context, user *model.User) error {
// 	return nil
// }

// func (r userRepository) CreateUserAddress(ctx context.Context, userAddress *model.UserAddress) error {
// 	e := newUserAddressEntity(userAddress)
// 	mutation := (*yo.UserAddress)(e).Insert(ctx)
// 	if _, err := r.ApplyMutations(ctx, []*spanner.Mutation{mutation}); err != nil {
// 		return err
// 	}

// 	return nil
// }
