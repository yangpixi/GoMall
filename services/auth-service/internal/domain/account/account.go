package account

import (
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

// New create a new account aggregate root instance
func New(id uint, username, password string, roleIDs []uint) (*Account, error) {
	if username == "" {
		return nil, ErrMissingUsername
	}

	if password == "" {
		return nil, ErrMissingPassword
	}

	if len(roleIDs) == 0 {
		return nil, ErrMissingRoles
	}

	return &Account{
		id:       id,
		username: username,
		password: password,
		status:   1,
		roleIDs:  roleIDs,
	}, nil

}

// Snapshot gives a snapshot for an account aggregate root,
// which provides data for persistence layer
func (a *Account) Snapshot() (*State, error) {
	return &State{
		ID:       a.id,
		Username: a.username,
		Password: a.password,
		Status:   1,
		RoleIDs:  a.roleIDs,
	}, nil
}

func (a *Account) ID() uint {
	return a.id
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
