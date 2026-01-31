package repo

import (
	"context"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/google/uuid"
)


type ProfilesRepo interface {
	CreateProfile(context.Context, entities.ProfileEntity) (*entities.ProfileEntity, error)
	UpdateProfile(context.Context, map[string]any) (*entities.ProfileEntity, error)
	GetProfileIdByEmail(context.Context, string) (uuid.UUID, error)
	AddPreferences(context.Context, entities.ProfileEntity) 
}