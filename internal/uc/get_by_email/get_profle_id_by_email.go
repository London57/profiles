package get_by_email

import (
	"context"

	"github.com/London57/profiles/internal/interfaces/repo"
	"github.com/google/uuid"
)

type GetProfileByEmail struct {
	repo repo.ProfilesRepo
}

func (uc GetProfileByEmail) Exec(ctx context.Context, email string) (uuid.UUID, error) {
	id, err := uc.repo.GetProfileIdByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, err
	}

	return id, err
}