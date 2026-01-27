package request

import (
	"time"

	"github.com/google/uuid"
)

type ProfileUpdateRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
	Username string `json:"username"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Birthday  time.Time `json:"birthday"`
	Name      string  `json:"name" binding:"max=30,omitempty"`
}