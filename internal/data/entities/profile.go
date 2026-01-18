package entities

import (
	"fmt"

	"github.com/London57/profiles/internal/consts"
	"github.com/google/uuid"
)

type ProfileEntity struct {
	ID uuid.UUID
	Name string
	Age int8
	Gender consts.Gender
	Longitude float32 
	Latitude float32
}

func (p ProfileEntity) Fields() (string, string) {
	return fmt.Sprintf(""), fmt.Sprintf()
}

func (ProfileEntity) New(id uuid.UUID, name string, age int8, gender consts.Gender, long float32, lat float32) ProfileEntity {
	return ProfileEntity{
		ID: id,
		Name: name, 
		Age: age,
		Gender: gender,
		Longitude: long,
		Latitude: lat,
	}
}