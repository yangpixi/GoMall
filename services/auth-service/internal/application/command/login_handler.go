package command

import (
	"context"
	"errors"
	"strconv"

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

	t, expire, err := h.issuer.Token(strconv.FormatInt(user.ID(), 10), user.RoleIDs())
	if err != nil {
		return nil, err
	}

	rt, err := h.issuer.RefreshToken(strconv.FormatInt(user.ID(), 10), user.RoleIDs())
	if err != nil {
		return nil, err
	}

	res := &LoginResult{
		Token:        t,
		RefreshToken: rt,
		Username:     user.Username(),
		ExpiresIn:    expire,
		TokenType:    "Bearer",
	}

	return res, nil
}
