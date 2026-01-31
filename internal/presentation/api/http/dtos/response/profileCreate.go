package response

import (
	"github.com/google/uuid"
)

type ProfileCreateResponse struct {
	ID uuid.UUID `json:"id"`
	Jwt_access_token string `json:"access_token"`
	Jwt_refresh_token string `json:"refresh_token"`
}