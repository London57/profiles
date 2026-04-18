package response

import "time"

type AddPreferencesResponse struct {
	Birthday *time.Time `json:"birthday,omitempty"`
	Raduis   *int16     `json:"radius,omitempty"`
}