package profile

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type UpdateProfileHandler struct {
	repo profile.Repository
}

func NewUpdateProfileHandler(repo profile.Repository) (*UpdateProfileHandler, error) {
	if repo == nil {
		return nil, errors.New("invalid profile repository")
	}

	return &UpdateProfileHandler{repo: repo}, nil
}

// Handle user profile updating command except user's avatar
func (h *UpdateProfileHandler) Handle(ctx context.Context, cmd *UpdateProfileCommand) error {
	userID, ok := id.UserIDFromCtx(ctx)

	if !ok {
		return errors.New("failed to retrieve userID from context")
	}

	p, err := h.repo.FindByUserID(ctx, userID)

	if err != nil {
		return err
	}

	if cmd.Phone != nil {
		if err := p.ChangePhone(*cmd.Phone); err != nil {
			return err
		}
	}

	if cmd.Email != nil {
		if err := p.ChangeEmail(*cmd.Email); err != nil {
			return err
		}
	}

	if cmd.Nickname != nil {
		if err := p.ChangeNickname(*cmd.Nickname); err != nil {
			return err
		}
	}

	err = h.repo.Update(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
