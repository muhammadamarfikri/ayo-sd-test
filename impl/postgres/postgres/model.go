package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	Name                string    `json:"name" db:"name"`
	LogoURL             *string   `json:"logo_url,omitempty" db:"logo_url"`
	FoundedYear         int       `json:"founded_year" db:"founded_year"`
	HeadquartersAddress string    `json:"headquarters_address" db:"headquarters_address"`
	HeadquartersCity    string    `json:"headquarters_city" db:"headquarters_city"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

type Player struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	TeamID      uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	Name        string    `json:"name" db:"name"`
	HeightCM    int       `json:"height_cm" db:"height_cm"`
	WeightKG    int       `json:"weight_kg" db:"weight_kg"`
	Position    string    `json:"position" db:"position"`
	ShirtNumber int       `json:"shirt_number" db:"shirt_number"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Matches struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	MatchAt    time.Time `json:"match_at"`
	HomeTeamID uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	AwayTeamID uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type MatchResult struct {
	MatchID    uuid.UUID `gorm:"type:uuid;default:uuidv7();primaryKey;column:match_id"`
	HomeScore  int       `json:"home_score" db:"home_score"`
	AwayScore  int       `json:"away_score" db:"away_score"`
	ReportedAt time.Time `json:"reported_at" db:"reported_at"`
	DeletedAt  time.Time `json:"deleted_at" db:"reported_at"`
}

type MatchGoal struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	MatchID   uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	PlayerID  uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	Minute    int       `json:"minute" db:"minute"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
