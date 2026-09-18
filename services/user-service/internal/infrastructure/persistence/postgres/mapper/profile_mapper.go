package mapper

import (
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
)

func ToProfile(p *model.Profile, addressIDs []int64) (*profile.Profile, error) {

	if p == nil {
		return nil, errors.New("invalid profile model")
	}

	return profile.Restore(&profile.State{
		UserID:     p.UserID,
		Nickname:   p.Nickname,
		Phone:      p.Phone,
		Email:      p.Email,
		Avatar:     p.Avatar,
		AddressIDs: addressIDs,
	})
}

func ToProfilePO(p *profile.Profile) (*model.Profile, error) {
	s, err := p.Snapshot()
	if err != nil {
		return nil, fmt.Errorf("failed to get a snapshot for profile: %w", err)
	}

	return &model.Profile{
		UserID:   s.UserID,
		Nickname: s.Nickname,
		Phone:    s.Phone,
		Email:    s.Email,
		Avatar:   s.Avatar,
	}, nil
}
