package profile

import "errors"

type State struct {
	UserID     uint
	Nickname   string
	Phone      string
	Email      string
	Avatar     string
	AddressIDs []uint
}

func Restore(s *State) (*Profile, error) {
	if s == nil || s.AddressIDs == nil {
		return nil, errors.New("invalid arguments")
	}

	if s.UserID == 0 {
		return nil, errors.New("invalid profile object")
	}

	return &Profile{
		userID:     s.UserID,
		nickname:   s.Nickname,
		phone:      s.Phone,
		email:      s.Email,
		avatar:     s.Email,
		addressIDs: s.AddressIDs,
	}, nil
}
