package entity

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID         uuid.UUID `json:"id" db:"id"`
	MatchAt    time.Time `json:"match_at" db:"match_at"`
	HomeTeamID uuid.UUID `json:"home_team_id" db:"home_team_id"`
	AwayTeamID uuid.UUID `json:"away_team_id" db:"away_team_id"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type MatchResult struct {
	MatchID    uuid.UUID `json:"match_id" db:"match_id"`
	HomeScore  int       `json:"home_score" db:"home_score"`
	AwayScore  int       `json:"away_score" db:"away_score"`
	ReportedAt time.Time `json:"reported_at" db:"reported_at"`
}

type MatchGoal struct {
	ID        uuid.UUID `json:"id" db:"id"`
	MatchID   uuid.UUID `json:"match_id" db:"match_id"`
	PlayerID  uuid.UUID `json:"player_id" db:"player_id"`
	Minute    int       `json:"minute" db:"minute"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
