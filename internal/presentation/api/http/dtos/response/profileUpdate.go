package response

import (
	"time"
)

type ProfileUpdateResponse struct {
	Latitude  *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Username  *string `json:"username,omitempty"`
	Name      *string  `json:"name,omitempty"`
}