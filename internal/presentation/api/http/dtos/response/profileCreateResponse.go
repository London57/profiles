package response

import (
	"github.com/London57/profiles/internal/consts"
	"github.com/google/uuid"
)

type ProfileCreateRequest struct {
	ID uuid.UUID `json:"id"`
	Latutude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Gender consts.Gender `json:"gender"`
	Age int8 `json:"age"`
	Name string `json:"name"`
}