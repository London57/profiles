package repo

import (
	"context"

	"github.com/London57/profiles/internal/data/entities"
)

type ProfilesRepo interface {
	CreateProfile(context.Context, entities.ProfileEntity) (entities.ProfileEntity, error)
	AddPreferences(context.Context, entities.ProfileEntity) 
}