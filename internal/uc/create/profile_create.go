package create

import (
	"context"
	"fmt"

	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/interfaces/repo"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	jwt "github.com/London57/profiles/pkg/jwtutil"
	"github.com/London57/profiles/pkg/password"
	"github.com/google/uuid"
)

type ProfileCreate struct {
	repo repo.ProfilesRepo
	jwtConfig config.JWT
	jwtutil jwt.JWT
}

func (ProfileCreate) NewProfleCreate(repo repo.ProfilesRepo, config config.JWT, jwtutil jwt.JWT) ProfileCreate {
	return ProfileCreate{
		jwtConfig: config,
		repo: repo,
		jwtutil: jwtutil,
	}
}

func (uc ProfileCreate) Exec(ctx context.Context, r request.ProfileCreateRequest) (response.ProfileCreateResponse, error) {
	pswd := r.Password
	pswd, err := password.GeneratePasswordHash(pswd)
	if err != nil {
		return response.ProfileCreateResponse{}, err
	}
	
	data := entities.ProfileEntity{
		ID: uuid.Nil,
		Name: r.Name,
		Birthday: r.Birthday,
		Email: r.Email,
		Username: r.Username,
		Password: pswd,
		Gender: r.Gender,
		Longitude: r.Longitude,
		Latitude: r.Latitude,
	}
	
	entity, err := uc.repo.CreateProfile(ctx, data)
	if err != nil {
		return response.ProfileCreateResponse{}, err
	}

	access_token, err := uc.jwtutil.CreateAccessToken(entity.ID, entity.Username, uc.jwtConfig.AccessTokenSecret, uc.jwtConfig.AccessTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateResponse{}, fmt.Errorf("jwt error: %w", err)
	}

	refresh_token, err := uc.jwtutil.CreateRefreshToken(entity.ID, entity.Name, uc.jwtConfig.RefreshTokenSecret, uc.jwtConfig.RefreshTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateResponse{}, fmt.Errorf("jwt error: %w", err)
	}

	resp := response.ProfileCreateResponse{
		ID: entity.ID,
		Jwt_access_token: access_token,
		Jwt_refresh_token: refresh_token,
	}
	return resp, err
}