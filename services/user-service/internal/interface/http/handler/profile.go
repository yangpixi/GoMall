package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/profile"
	"github.com/yangpixi/GoMall/shared/response"
)

type ProfileHandler struct {
	createHandler *profile.CreateProfileHandler
	updateHandler *profile.UpdateProfileHandler
}

type createProfileRequest struct {
	UserID   int64  `json:"userId,string"` // not required for user, but required for admin
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type updateProfileRequest struct {
	UserID   int64   `json:"userId,string"`
	Nickname *string `json:"nickname"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}

func NewProfileHandler(ch *profile.CreateProfileHandler, uh *profile.UpdateProfileHandler) (*ProfileHandler, error) {
	if ch == nil || uh == nil {
		return nil, errors.New("invalid profile handler")
	}
	return &ProfileHandler{createHandler: ch, updateHandler: uh}, nil
}

// CreateHandler handle profile creation request
func (h *ProfileHandler) CreateHandler(c *gin.Context) {
	var req createProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.createHandler.Handle(c.Request.Context(), &profile.CreateProfileCommand{
		UserID:   req.UserID,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK[string](c, "profile creation successfully")
}

func (h *ProfileHandler) UpdateHandler(c *gin.Context) {
	var req updateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.updateHandler.Handle(c.Request.Context(), &profile.UpdateProfileCommand{
		UserID:   req.UserID,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK[string](c, "profile updating successfully")
}
