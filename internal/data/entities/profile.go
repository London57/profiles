package entities

import (
	"time"

	"github.com/London57/profiles/internal/consts"
	"github.com/google/uuid"
)

type ProfileEntity struct {
	ID uuid.UUID
	Email string
	Username string
	Phone_number string
	Password string
	Name string
	Birthday time.Time
	Gender consts.Gender
	Longitude float32 
	Latitude float32
}

func (ProfileEntity) New(id uuid.UUID, email, name, username, password, phone_number string, birthday time.Time, gender consts.Gender, long float32, lat float32) ProfileEntity {

	return ProfileEntity{
		ID: id,
		Email: email,
		Name: name, 
		Phone_number: phone_number,
		Birthday: birthday,
		Password: password,
		Gender: gender,
		Longitude: long,
		Latitude: lat,
	}
}