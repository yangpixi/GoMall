package command

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/account"
)

type LoginHandler struct {
	repo   account.Repository
	issuer TokenIssuer
}

func NewLoginHandler(repo account.Repository, issue TokenIssuer) (*LoginHandler, error) {
	if repo == nil || issue == nil {
		return nil, errors.New("missing necessary arguments")
	}

	return &LoginHandler{repo: repo, issuer: issue}, nil
}

// Handle login process
func (h *LoginHandler) Handle(ctx context.Context, cmd *LoginCommand) (*LoginResult, error) {
	user, err := h.repo.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}

	if err = user.CanLoginIn(); err != nil {
		return nil, err
	}

	if err = user.Validate(cmd.Password); err != nil {
		return nil, err
	}

	t, err := h.issuer.Token(user.Username(), user.RoleIDs())
	if err != nil {
		return nil, err
	}

	rt, err := h.issuer.RefreshToken(user.Username(), user.RoleIDs())
	if err != nil {
		return nil, err
	}

	res := &LoginResult{
		Token:        t,
		RefreshToken: rt,
	}

	return res, nil
}
