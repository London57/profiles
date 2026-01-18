package request

import "github.com/London57/profiles/internal/consts"

type ProfileCreateRequest struct {
	Latutude float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
	Gender consts.Gender `json:"gender" binding:"required"`
	Age int8 `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
}