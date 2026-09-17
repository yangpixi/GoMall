package profile

type Profile struct {
	userID     int64
	nickname   string
	phone      string
	email      string
	avatar     string
	addressIDs []int64
}

// New creates a new profile for a user.
// It is usually called during registration, so the address
// can be left unset and filled in later.
func New(userID int64, nickname, phone, email, avatar string) (*Profile, error) {
	if nickname == "" || phone == "" || userID == 0 {
		return nil, ErrInvalidProfile
	}

	return &Profile{
		userID:   userID,
		nickname: nickname,
		phone:    phone,
		email:    email,
		avatar:   avatar,
	}, nil

}

// AppendAddress adds a new address for a user
func (p *Profile) AppendAddress(addressID int64) error {
	if addressID == 0 {
		return ErrInvalidAddressID
	}
	p.addressIDs = append(p.addressIDs, addressID)
	return nil
}
