package request

import (
	"time"

	"github.com/google/uuid"
)

type AddPreferencesRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
	Birthday time.Time `json:"birthday"`
	Raduis int16 `json:"radius"`
}