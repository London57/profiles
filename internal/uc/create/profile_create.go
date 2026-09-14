package uc_create

import (
	"context"
	"fmt"

	jwt "github.com/London57/jwt-auth/jwtutil"
	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/datagen"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	"github.com/London57/profiles/pkg/password"
)

type repo interface {
	CreateProfile(context.Context, datagen.CreateProfileParams) (datagen.SocialProfile, error)
}

type ProfileCreate struct {
	repo repo
	jwtConfig config.JwtConfig
	jwtutil jwt.Jwt
}

func (ProfileCreate) New(repo repo, config config.JwtConfig, jwtutil jwt.Jwt) ProfileCreate {
	return ProfileCreate{
		jwtConfig: config,
		repo: repo,
		jwtutil: jwtutil,
	}
}

func (uc ProfileCreate) Exec(ctx context.Context, data datagen.CreateProfileParams) (response.ProfileCreateResponse, error) {
	hash, err := password.GeneratePasswordHash(data.Password)
	if err != nil {
		return response.ProfileCreateResponse{}, err
	}
	data.Password = hash
	
	profile, err := uc.repo.CreateProfile(ctx, data)
	if err != nil {
		return response.ProfileCreateResponse{}, err
	}

	access_token, err := uc.jwtutil.CreateAccessToken(profile.ID, profile.Username, uc.jwtConfig.AccessTokenSecret, uc.jwtConfig.AccessTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateResponse{}, fmt.Errorf("jwt error: %w", err)
	}

	refresh_token, err := uc.jwtutil.CreateRefreshToken(profile.ID, profile.Name, uc.jwtConfig.RefreshTokenSecret, uc.jwtConfig.RefreshTokenExpiryHour)
	if err != nil {
		return response.ProfileCreateResponse{}, fmt.Errorf("jwt error: %w", err)
	}

	resp := response.ProfileCreateResponse{
		ID: profile.ID,
		Jwt_access_token: access_token,
		Jwt_refresh_token: refresh_token,
	}
	return resp, err
}