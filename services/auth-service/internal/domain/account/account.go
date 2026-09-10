package account

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Status int8

var (
	Enabled  Status = 1
	Disabled Status = 0
)

type Account struct {
	id       uint
	username string
	password string
	status   Status
	roleIDs  []uint
}

func New(username, password string) (*Account, error) {
	if username == "" {
		return nil, errors.New("username can not be nil")
	}

	if password == "" {
		return nil, errors.New("password can not be nil")
	}

	return &Account{
		username: username,
		password: password,
		status:   0,
	}, nil

}

func (a *Account) Username() string {
	return a.username
}
func (a *Account) RoleIDs() []uint {
	return a.roleIDs
}

// CanLoginIn check if a user is banned
func (a *Account) CanLoginIn() error {
	if a.status != 1 {
		return ErrBannedAccount
	}
	return nil
}

// Validate if password passed in equals to password saved in database
func (a *Account) Validate(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(a.password), []byte(password)); err != nil {
		return ErrUsernamePasswordIncorrect
	}

	return nil
}
