package profile

type Profile struct {
	userID     uint
	nickname   string
	phone      string
	email      string
	avatar     string
	addressIDs []uint
}

func New(userID uint, nickname, phone, email, avatar string, addressIDs []uint) (*Profile, error) {
	if nickname == "" || phone == "" || userID == 0 {
		return nil, ErrInvalidProfile
	}

	return &Profile{
		userID:     userID,
		nickname:   nickname,
		phone:      phone,
		email:      email,
		avatar:     avatar,
		addressIDs: addressIDs,
	}, nil

}
