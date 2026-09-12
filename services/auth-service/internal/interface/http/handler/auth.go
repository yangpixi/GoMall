package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/auth-service/internal/application/command"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

type AuthHandler struct {
	login    *command.LoginHandler
	register *command.RegisterHandler
	refresh  *command.RefreshHandler
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type refreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func NewAuthHandler(login *command.LoginHandler, register *command.RegisterHandler, refresh *command.RefreshHandler) (*AuthHandler, error) {
	if login == nil || register == nil || refresh == nil {
		return nil, errors.New("missing required arguments")
	}

	return &AuthHandler{login: login, register: register, refresh: refresh}, nil
}

func (a *AuthHandler) LoginHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	result, err := a.login.Handle(c.Request.Context(), &command.LoginCommand{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[*command.LoginResult](c, result)

}

func (a *AuthHandler) RegisterHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	err := a.register.Handle(c.Request.Context(), &command.RegisterCommand{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[string](c, "register successfully")
}

func (a *AuthHandler) RefreshHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	res, err := a.refresh.Handle(c.Request.Context(), &command.RefreshCommand{RefreshToken: req.RefreshToken})

	if err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[*command.RefreshResult](c, res)
}
