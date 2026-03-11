package get_by_email

import (
	"context"

	"github.com/google/uuid"
)

type repo interface {
	GetProfileIdByEmail(context.Context, string) (uuid.UUID, error)
}

type GetProfileByEmail struct {
	repo repo
}

func (GetProfileByEmail) NewGetProfileByEmail(repo repo) GetProfileByEmail {
	return GetProfileByEmail{
		repo: repo,
	}
}

func (uc GetProfileByEmail) Exec(ctx context.Context, email string) (uuid.UUID, error) {
	id, err := uc.repo.GetProfileIdByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, err
	}

	return id, err
}