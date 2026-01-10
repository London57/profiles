package repo

import (
	"context"
	"github.com/London57/profiles/internal/dtos"	
)

type ProfileRepo interface {
	CreateProfile(context.Context, dtos.ProfileCreateRequest)
	AddPreferences(context.Context, dtos.AddPreferencesRequest)
}