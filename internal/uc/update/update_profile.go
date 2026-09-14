package uc_update

import (
	"context"

	"github.com/London57/profiles/internal/data/datagen"
)

type repo interface {
	UpdateProfile(context.Context, datagen.UpdateProfileParams) (datagen.SocialProfile, error)
}

type ProfileUpdate struct {
	repo repo
}

func (ProfileUpdate) New(repo repo) ProfileUpdate {
	return ProfileUpdate{
		repo: repo,
	}
}

func (uc ProfileUpdate) Exec(ctx context.Context, data datagen.UpdateProfileParams) (datagen.SocialProfile, error) {
	profile, err := uc.repo.UpdateProfile(ctx, data)
	if err != nil {
		return datagen.SocialProfile{}, err
	}

	return profile, nil
}
