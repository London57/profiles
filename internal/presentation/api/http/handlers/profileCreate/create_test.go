package create

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/entities"
	repo "github.com/London57/profiles/internal/interfaces/repo/mocks"
	"github.com/London57/profiles/internal/uc/create"
	"github.com/London57/profiles/internal/uc/get_by_email"
	mock_jwtutil "github.com/London57/profiles/pkg/jwtutil/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)


func TestCreate(t *testing.T) {
	uid, _ := uuid.NewRandom()
	tcs := []struct{
		name          string
		requestBody	  string
		expectResponseBody string
		expectedStatusCode int
		setup func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT)
	}{
		{
			name: "success",
			requestBody: `{"longitude": 16.0,
				"latitude": 14.01,
				"email": "dad@adad.ru",
				"username": "123dadas",
				"gender": 0,
				"birthday": "1998-06-13T15:04:05Z",
				"name": "123sdfs",
				"password": "123131"}`,
			expectResponseBody: fmt.Sprintf(`{"id":"%s","access_token":"12313","refresh_token":"12313"}`, uid),
			expectedStatusCode: 201,
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT) {
				repo := repo.NewMockProfilesRepo(ctrl)
				jwt := mock_jwtutil.NewMockJWT(ctrl)

				jwt.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("12313", nil)
				jwt.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("12313", nil)
				repo.EXPECT().GetProfileIdByEmail(gomock.Any(), gomock.Any()).Return(uuid.Nil, pgx.ErrNoRows)
				repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(&entities.ProfileEntity{
					ID: uid,
					Latitude: 14.01,
					Longitude: 16,
					Email: "dad@adad.ru",
					Username: "123dadas",
					Gender: 0,
					Birthday: time.Date(1998, 6, 13, 3, 2, 0, 0, time.UTC),
					Name: "123sdfs",
					Password: "",
					}, nil)
				return repo, jwt
			},
		},
		{
			name: "gender field is empty error",
			requestBody: `{"longitude": 16.0,
				"latitude": 14.01,
				"email": "dad@adad.ru",
				"username": "123dadas",
				"birthday": "1998-06-13T15:04:05Z",
				"name": "123sdfs",
				"password": "123131"}`,
			expectResponseBody: `{"error":"failed to parse JSON: Key: 'ProfileCreateRequest.Gender' Error:Field validation for 'Gender' failed on the 'required' tag"}`,
			expectedStatusCode: 400,
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT) {
				repo := repo.NewMockProfilesRepo(ctrl)
				jwt := mock_jwtutil.NewMockJWT(ctrl)
				return repo, jwt
			},
		},
		{
			name: "user with this email already exists",
			requestBody: `{"longitude": 16.0,
				"latitude": 14.01,
				"email": "dad@adad.ru",
				"username": "123dadas",
				"gender": 0,
				"birthday": "1998-06-13T15:04:05Z",
				"name": "123sdfs",
				"password": "123131"}`,
			expectResponseBody: `{"error":"user with this email already exists"}`,
			expectedStatusCode: 409,
			setup: func(ctrl *gomock.Controller) (*repo.MockProfilesRepo, *mock_jwtutil.MockJWT) {
				repo := repo.NewMockProfilesRepo(ctrl)
				jwt := mock_jwtutil.NewMockJWT(ctrl)
				repo.EXPECT().GetProfileIdByEmail(gomock.Any(), gomock.Any()).Return(uid, nil)
				return repo, jwt
			},
		}, 
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			repo, jwt := tc.setup(ctrl)

			create_uc := create.ProfileCreate{}.NewProfleCreate(repo, config.JWT{}, jwt)
			get_by_email_uc := get_by_email.GetProfileByEmail{}.NewGetProfileByEmail(repo)
			handler := ProfileCreateHandler{}.NewProfleCreateHandler(create_uc, get_by_email_uc)

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(
				"POST",
				"/profiles/registration",
				strings.NewReader(tc.requestBody),
			)
			c.Request.Header.Set("Content-Type", "application/json")
			
			
			handler.CreateProfile(c)
			defer c.Request.Body.Close()
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectResponseBody, w.Body.String())
			
		})
	}
}