package interactors

import (
	"context"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/interfaces/repo"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/request"
	"github.com/London57/profiles/internal/presentation/api/http/dtos/response"
	"github.com/google/uuid"
)

type ProfileCreate struct {
	repo repo.ProfilesRepo
}

func (uc ProfileCreate) Exec(ctx context.Context, r request.ProfileCreateRequest) (response.ProfileCreateRequest, error) {
	data := entities.ProfileEntity{
		ID: uuid.Nil,
		Name: r.Name,
		Age: r.Age,
		Gender: r.Gender,
		Longitude: r.Longitude,
		Latitude: r.Latutude,
	}
	
	entity, err := uc.repo.CreateProfile(ctx, data)
	if err != nil {
		return response.ProfileCreateRequest{}, err
	}
	resp := response.ProfileCreateRequest{
		ID: entity.ID,
		Name: entity.Name,
		Age: entity.Age,
		Gender: entity.Gender,
		Latutude: entity.Latitude,
		Longitude: entity.Longitude,
	}
	return resp, err
}