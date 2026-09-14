package response

import (
	"gopkg.in/guregu/null.v4"
)

type AddPreferencesResponse struct {
	AgeTo null.Int `json:"age_to"` 
	AgeFrom null.Int `json:"age_from"` 
	Raduis   null.Int     `json:"radius"`
} 	