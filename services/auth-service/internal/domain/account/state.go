package account

import (
	"errors"
)

type State struct {
	ID       uint
	Username string
	Password string
	Status   Status
	RoleIDs  []uint
}

func RestoreAccount(s *State) (*Account, error) {

	if s.Username == "" || s.Password == "" {
		return nil, errors.New("invalid account state")
	}

	return &Account{
		id:       s.ID,
		username: s.Username,
		password: s.Password,
		status:   s.Status,
		roleIDs:  s.RoleIDs,
	}, nil
}
