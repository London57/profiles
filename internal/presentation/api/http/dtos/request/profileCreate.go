package request

import (
	"time"

	"github.com/London57/profiles/internal/consts"
	"gopkg.in/guregu/null.v4"
)

type ProfileCreateRequest struct {
	Latitude null.Float `json:"latitude"`
	Longitude null.Float `json:"longitude"`
	Email string `json:"email" binding:"required,email"`
	Phone_number null.String `json:"phone_number"`
	Username string `json:"username" binding:"required,min=6,max=30"`
	Gender consts.Gender `json:"gender" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
	Name string `json:"name" binding:"required,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}