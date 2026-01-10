package dtos

import (
	"github.com/google/uuid"
)

type AddPreferencesRequest struct {
	id uuid.UUID `json:"id" binding:"required"`
	Age int8 `json:"age" binding:"required"`
	Raduis *int16 `json:"radius" binding:"omitempty"`
}