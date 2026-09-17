package role

import (
	"errors"
)

type State struct {
	ID            int64
	Code          string
	Name          string
	PermissionIDs []int64
}

func Restore(s *State) (*Role, error) {
	if s == nil {
		return nil, errors.New("invalid role state")
	}

	return &Role{
		id:            s.ID,
		code:          s.Code,
		name:          s.Name,
		permissionIDs: s.PermissionIDs,
	}, nil
}
