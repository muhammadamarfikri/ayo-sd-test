package entity

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID          uuid.UUID `json:"id" db:"id"`
	TeamID      uuid.UUID `json:"team_id" db:"team_id"`
	Name        string    `json:"name" db:"name"`
	HeightCM    int       `json:"height_cm" db:"height_cm"`
	WeightKG    int       `json:"weight_kg" db:"weight_kg"`
	Position    string    `json:"position" db:"position"`
	ShirtNumber int       `json:"shirt_number" db:"shirt_number"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
