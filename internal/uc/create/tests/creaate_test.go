package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/entities"
	repo "github.com/London57/profiles/internal/interfaces/repo/mocks"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	interactors "github.com/London57/profiles/internal/uc/create"
	mock_jwtutil "github.com/London57/profiles/pkg/jwtutil/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func TestCreate(t *testing.T) {
	tcs := []struct{
		name string
		req request.ProfileCreateRequest
		expected_resp response.ProfileCreateResponse
        setup func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT)
		expected_err error
	}{
		{
			name: "success",
			req: request.ProfileCreateRequest{
				Latitude: 14.01,
				Longitude: 16,
				Email: "dad@adad.ru",
				Username: "123",
				Gender: 0,
				Birthday: time.Date(1998, 6, 13, 3, 2, 0, 0, time.UTC),
				Name: "123sdfs",
				Password: "12313",
			},
			expected_resp: response.ProfileCreateResponse{
				ID: uuid.UUID{},
				Jwt_access_token: "12313",
				Jwt_refresh_token: "12313",
			},
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT) {
				repo := repo.NewMockProfilesRepo(ctrl)
				jwt := mock_jwtutil.NewMockJWT(ctrl)

				jwt.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("12313", nil)
				jwt.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("12313", nil)
				repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(&entities.ProfileEntity{}, nil)
				return repo, jwt
			},
			expected_err: nil,
		},
{
			name: "jwt error",
			req: request.ProfileCreateRequest{
				Latitude: 14.01,
				Longitude: 16,
				Email: "dad@adad.ru",
				Username: "123",
				Gender: 0,
				Birthday: time.Date(1998, 6, 13, 3, 2, 0, 0, time.UTC),
				Name: "123sdfs",
				Password: "12313",
			},
			expected_resp: response.ProfileCreateResponse{
				ID: uuid.UUID{},
				Jwt_access_token: "",
				Jwt_refresh_token: "",
			},
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT) {
				repo := repo.NewMockProfilesRepo(ctrl)
				jwt := mock_jwtutil.NewMockJWT(ctrl)

				jwt.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", errors.New("failed to sign jwt"))
				repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(&entities.ProfileEntity{}, nil)
				return repo, jwt
			},
			expected_err: fmt.Errorf("jwt error: %w", errors.New("failed to sign jwt")),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo, jwt := tc.setup(ctrl)
			uc := interactors.ProfileCreate{}.NewProfleCreate(repo, config.JWT{}, jwt)

			ctx := context.Background()
			res, err := uc.Exec(ctx, tc.req)
			assert.Equal(t, res, tc.expected_resp)
			
			assert.Equal(t, err, tc.expected_err)
		})
	}
}