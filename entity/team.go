package entity

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	Name                string    `json:"name" db:"name"`
	LogoURL             *string   `json:"logo_url,omitempty" db:"logo_url"`
	FoundedYear         int       `json:"founded_year" db:"founded_year"`
	HeadquartersAddress string    `json:"headquarters_address" db:"headquarters_address"`
	HeadquartersCity    string    `json:"headquarters_city" db:"headquarters_city"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
