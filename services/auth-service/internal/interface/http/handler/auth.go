package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/auth-service/internal/application/command"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

type AuthHandler struct {
	login    *command.LoginHandler
	register *command.RegisterHandler
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(login *command.LoginHandler, register *command.RegisterHandler) (*AuthHandler, error) {
	if login == nil || register == nil {
		return nil, errors.New("missing required arguments")
	}

	return &AuthHandler{login: login, register: register}, nil
}

func (a *AuthHandler) LoginHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			response.Fail(c, bizErr, http.StatusOK)
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
			response.Fail(c, bizErr, http.StatusOK)
			c.Abort()
			return
		}
	}

	response.OK[*command.LoginResult](c, result)

}

func (a *AuthHandler) RegisterHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			response.Fail(c, bizErr, http.StatusOK)
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
			response.Fail(c, bizErr, http.StatusOK)
			c.Abort()
			return
		}
	}

	response.OK[string](c, "register successfully")
}
