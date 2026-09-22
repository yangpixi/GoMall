package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/profile"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

type ProfileHandler struct {
	handler *profile.CreateProfileHandler
}

type createProfileRequest struct {
	UserID   int64  `json:"userId" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func NewProfileHandler(h *profile.CreateProfileHandler) (*ProfileHandler, error) {
	if h == nil {
		return nil, errors.New("invalid profile handler")
	}
	return &ProfileHandler{handler: h}, nil
}

// CreateHandler handle profile creation request
func (h *ProfileHandler) CreateHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req createProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	err := h.handler.Handle(c.Request.Context(), &profile.CreateProfileCommand{
		UserID:   req.UserID,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Avatar,
	})

	if err = c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[string](c, "profile creation successfully")
}
