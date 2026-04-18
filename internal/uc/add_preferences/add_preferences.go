package add_preferences

import (
	"context"
	"time"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
)

type repo interface {
	AddPreferences(context.Context, map[string]any) (entities.Preferences, error)
}

type AddPreferences struct {
	repo repo
}

func (AddPreferences) New(repo repo) AddPreferences{
	return AddPreferences{
		repo: repo,
	}
}

func (uc AddPreferences) Exec(ctx context.Context, req request.AddPreferencesRequest) (response.AddPreferencesResponse, error) {
	fields := make(map[string]any, 2)

	fields["profile_id"] = req.ID

	birthday := req.Birthday
	if birthday != (time.Time{}) {
		fields["birthday"] = birthday
	}
	
	raduis := req.Raduis
	if raduis != 0 {
		fields["radius"] = raduis
	}

	preferences, err := uc.repo.AddPreferences(ctx, fields)
	if err != nil {
		return response.AddPreferencesResponse{}, err
	}

	return response.AddPreferencesResponse{
		Birthday: &preferences.Birthday,
		Raduis: &preferences.Radius,
	}, nil
}