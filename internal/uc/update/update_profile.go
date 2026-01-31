package update

import (
	"context"
	"time"

	"github.com/London57/profiles/internal/interfaces/repo"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
)

type ProfileUpdate struct {
	repo repo.ProfilesRepo
}

func (ProfileUpdate) New(repo repo.ProfilesRepo) ProfileUpdate {
	return ProfileUpdate{
		repo: repo,
	}
}

func (uc ProfileUpdate) Exec(ctx context.Context, req request.ProfileUpdateRequest) (response.ProfileUpdateResponse, error) {
	
	fields := make(map[string]any, 2)

	username := req.Username
	if username != "" {
		fields["username"] = username
	}

	name := req.Name
	if name != "" {
		fields["name"] = name
	}

	birthday := req.Birthday
	if birthday != (time.Time{}) {
		fields["birthday"] = birthday
	}

	longitude := req.Longitude
	if longitude != 0 {
		fields["longitude"] = longitude
	}

	latitude := req.Latitude
	if latitude != 0 {
		fields["latitude"] = latitude
	}
	res, err := uc.repo.UpdateProfile(ctx, fields)
	if err != nil {
		return response.ProfileUpdateResponse{}, err
	}
	resp := response.ProfileUpdateResponse{
		Latitude: &res.Latitude,
		Longitude: &res.Longitude,
		Name: &res.Name,
		Birthday: &res.Birthday,
		Username: &res.Username,
	}

	return resp, nil
}