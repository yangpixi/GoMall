package profile

import (
	"regexp"
	"strings"
)

type Profile struct {
	userID     int64
	nickname   string
	phone      string
	email      string
	avatar     string
	addressIDs []int64
}

// phonePattern validates phone number user changed
var phonePattern = regexp.MustCompile(`^1[3-9][0-9]{9}$`)

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

func (p *Profile) Snapshot() (*State, error) {
	return &State{
		UserID:     p.userID,
		Nickname:   p.nickname,
		Phone:      p.phone,
		Email:      p.email,
		Avatar:     p.avatar,
		AddressIDs: p.addressIDs,
	}, nil
}

func (p *Profile) ChangeNickname(name string) error {
	if strings.TrimSpace(name) == "" || len(name) > 10 || len(name) < 3 {
		return ErrInvalidNickname
	}

	p.nickname = name
	return nil
}

func (p *Profile) ChangePhone(phone string) error {
	phone = strings.TrimSpace(phone)

	if !phonePattern.MatchString(phone) {
		return ErrInvalidPhone
	}

	p.phone = phone
	return nil
}

func (p *Profile) ChangeEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return ErrInvalidEmail
	}

	p.email = email
	return nil
}

// ChangeAvatar only changes avatar url from oss
func (p *Profile) ChangeAvatar(url string) error {
	if strings.TrimSpace(url) == "" {
		return ErrInvalidAvatar
	}

	p.avatar = url
	return nil
}
