package entities

import (
	"time"
)

type Preferences struct {
	Birthday   time.Time
	Longitude  float32
	Latitude   float32
}
