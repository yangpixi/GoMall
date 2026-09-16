package mapper

import (
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
)

func ToProfile(p *model.Profile, addressIDs []uint) (*profile.Profile, error) {

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
