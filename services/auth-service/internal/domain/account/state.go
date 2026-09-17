package account

import (
	"errors"
)

type State struct {
	ID       int64
	Username string
	Password string
	Status   Status
	RoleIDs  []int64
}

func Restore(s *State) (*Account, error) {

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
