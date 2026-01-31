package tests

import (
	"context"
	"testing"
	"time"

	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/entities"
	repo "github.com/London57/profiles/internal/interfaces/repo/mocks"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	interactors "github.com/London57/profiles/internal/uc/update"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func Ptr[t any](v t) *t {
	return &v
}

func TestUpdate(t *testing.T) {
	tcs := []struct{
		name          string
		req           request.ProfileUpdateRequest
		expected_resp response.ProfileUpdateResponse
		setup         func(ctrl *gomock.Controller) *repo.MockProfilesRepo
		expected_err  error
		
	}{
		{
			name: "success",
			req: request.ProfileUpdateRequest{
				Latitude:  14.01,
				Longitude: 16,
				Username:  "123",
				Birthday:  time.Date(1998, 6, 13, 3, 2, 0, 0, time.UTC),
				Name:      "123sdfs",
			},
			expected_resp: response.ProfileUpdateResponse{
				Latitude:  Ptr[float32](14.02),
				Longitude: Ptr[float32](16),
				Username:  Ptr[string]("123"),
				Birthday:  Ptr[time.Time](time.Date(1998, 6, 13, 3, 2, 0, 0, time.UTC)),
				Name:      Ptr[string]("123sdfs"),
			},
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo) {
				repo := repo.NewMockProfilesRepo(ctrl)

				repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(&entities.ProfileEntity{}, nil)
				return repo
			},
			expected_err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := tc.setup(ctrl)
			uc := interactors.ProfileUpdate{}.(repo, config.JWT{}, jwt)

			ctx := context.Background()
			res, err := uc.Exec(ctx, tc.req)
			assert.Equal(t, res, tc.expected_resp)

			assert.Equal(t, err, tc.expected_err)
		})
	}
}