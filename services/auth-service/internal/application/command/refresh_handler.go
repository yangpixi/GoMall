package command

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/shared/errs"
)

type RefreshHandler struct {
	issuer TokenIssuer
}

func NewRefreshHandler(issuer TokenIssuer) (*RefreshHandler, error) {
	if issuer == nil {
		return nil, errors.New("invalid token issuer")
	}
	return &RefreshHandler{issuer: issuer}, nil
}

func (h *RefreshHandler) Handle(_ context.Context, cmd *RefreshCommand) (*RefreshResult, error) {
	username, kind, roleIDs, err := h.issuer.ParseAndValidate(cmd.RefreshToken)
	if err != nil || kind != "refresh" {
		return nil, errs.New(400, "invalid token")
	}

	newToken, expiration, err := h.issuer.Token(username, roleIDs)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := h.issuer.RefreshToken(username, roleIDs)
	if err != nil {
		return nil, err
	}

	return &RefreshResult{
		Token:        newToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    expiration,
		TokenType:    "Bearer",
		Username:     username,
	}, nil
}
