package add_preferences

import (
	"context"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/google/uuid"
)

type repo interface {
	AddPreferences(context.Context, uuid.UUID, entities.Preferences) error
}

type AddPreferences struct {
	repo repo
}

func (AddPreferences) New(repo repo)