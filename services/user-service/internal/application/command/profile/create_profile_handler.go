package profile

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreateProfileHandler struct {
	repo profile.Repository
}

func NewCreateProfileHandler(repo profile.Repository) (*CreateProfileHandler, error) {
	if repo == nil {
		return nil, errors.New("invalid address repo")
	}
	return &CreateProfileHandler{repo: repo}, nil
}

func (h *CreateProfileHandler) Handle(ctx context.Context, cmd *CreateProfileCommand) error {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return profile.ErrInvalidUserID
	}

	p, err := profile.New(userID, cmd.Nickname, cmd.Phone, cmd.Email, cmd.Avatar)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
