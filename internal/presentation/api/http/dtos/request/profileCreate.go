package request

import (
	"time"

	"github.com/London57/profiles/internal/consts"
)

type ProfileCreateRequest struct {
	Latutude float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=6,max=30"`
	Gender consts.Gender `json:"gender" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
	Name string `json:"name" binding:"required,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}