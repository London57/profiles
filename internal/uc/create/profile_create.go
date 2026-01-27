package interactors

import (
	"context"
	"fmt"

	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/interfaces/repo"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	"github.com/London57/profiles/pkg/jwtutil"
	"github.com/London57/profiles/pkg/password"
	"github.com/google/uuid"
)

type ProfileCreate struct {
	repo repo.ProfilesRepo
	config.JWT
}

func (uc ProfileCreate) Exec(ctx context.Context, r request.ProfileCreateRequest) (response.ProfileCreateRequest, error) {
	pswd := r.Password
	pswd, err := password.GeneratePasswordHash(pswd)
	if err != nil {
		return response.ProfileCreateRequest{}, err
	}
	
	data := entities.ProfileEntity{
		ID: uuid.Nil,
		Name: r.Name,
		Birthday: r.Birthday,
		Password: pswd,
		Gender: r.Gender,
		Longitude: r.Longitude,
		Latitude: r.Latutude,
	}
	
	entity, err := uc.repo.CreateProfile(ctx, data)
	if err != nil {
		return response.ProfileCreateRequest{}, err
	}

	access_token, err := jwtutil.CreateAccessToken(entity.ID, entity.Username, uc.AccessTokenSecret, uc.AccessTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateRequest{}, fmt.Errorf("jwt error: %w", err)
	}

	refresh_token, err := jwtutil.CreateRefreshToken(entity.ID, entity.Name, uc.RefreshTokenSecret, uc.RefreshTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateRequest{}, fmt.Errorf("jwt error: %w", err)
	}

	resp := response.ProfileCreateRequest{
		ID: entity.ID,
		Jwt_access_token: access_token,
		Jwt_refresh_token: refresh_token,
	}
	return resp, err
}