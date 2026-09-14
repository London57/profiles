package uc_add_preferences

import (
	"context"

	"github.com/London57/profiles/internal/data/datagen"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
)

type repo interface {
	AddPreferences(context.Context, datagen.UpdatePreferencesParams) (datagen.SocialPreference, error)
}

type AddPreferences struct {
	repo repo
}

func (AddPreferences) New(repo repo) AddPreferences{
	return AddPreferences{
		repo: repo,
	}
}

func (uc AddPreferences) Exec(ctx context.Context, data datagen.UpdatePreferencesParams) (response.AddPreferencesResponse, error) {
	preferences, err := uc.repo.AddPreferences(ctx, data)
	if err != nil {
		return response.AddPreferencesResponse{}, err
	}

	return response.AddPreferencesResponse{
		AgeTo: preferences.AgeTo,
		AgeFrom: preferences.AgeFrom,
		Raduis: preferences.Raduis,
	}, nil
}