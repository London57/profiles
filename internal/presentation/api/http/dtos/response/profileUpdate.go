package response

import (
	"gopkg.in/guregu/null.v4"
)

type ProfileUpdateResponse struct {
	Latitude  null.Float `json:"latitude"`
	Longitude null.Float `json:"longitude"`
	Birthday  null.Float `json:"birthday"`
	Username  null.String `json:"username"`
	Name      null.String  `json:"name"`
	Phone_number null.String `json:"phone_number"`
}