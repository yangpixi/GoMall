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
	id       int64
	username string
	password string
	status   Status
	roleIDs  []int64
}

// New create a new account aggregate root instance
func New(id int64, username, password string, roleIDs []int64) (*Account, error) {
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

func (a *Account) ID() int64 {
	return a.id
}

func (a *Account) Username() string {
	return a.username
}

func (a *Account) RoleIDs() []int64 {
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
