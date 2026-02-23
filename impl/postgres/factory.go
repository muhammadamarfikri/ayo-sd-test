package postgres

import (
	"ayo/impl/postgres/internal"
	"ayo/impl/postgres/postgres"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreatTeam(ctx context.Context, teamModel *postgres.Team) (*postgres.Team, error)
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (*postgres.Team, error)
	UpdateTeamSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Team, error)
	DeleteTeam(ctx context.Context, id uuid.UUID) error
}

type PlayerRepository interface {
	CreatPlayer(ctx context.Context, playerModel *postgres.Player) (*postgres.Player, error)
	GetPlayerByID(ctx context.Context, playerID uuid.UUID) (*postgres.Player, error)
	UpdatePlayerSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Player, error)
	DeletePlayer(ctx context.Context, id uuid.UUID) error
}

type MatchRepository interface {
	CreatMatch(ctx context.Context, matchModel *postgres.Matches) (*postgres.Matches, error)
	GetMatchByID(ctx context.Context, matchID uuid.UUID) (*postgres.Matches, error)
	UpdateMatchSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Matches, error)
	DeleteMatch(ctx context.Context, id uuid.UUID) error
}

type MatchResultRepository interface {
	CreatMatchResult(ctx context.Context, matchModel *postgres.MatchResult) (*postgres.MatchResult, error)
	GetMatchResultByID(ctx context.Context, matchID uuid.UUID) (*postgres.MatchResult, error)
	UpdateMatchResultSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchResult, error)
	DeleteMatchResult(ctx context.Context, id uuid.UUID) error
}

type MatchGoalRepository interface {
	CreatMatchGoal(ctx context.Context, matchModel *postgres.MatchGoal) (*postgres.MatchGoal, error)
	GetMatchGoalByID(ctx context.Context, matchID uuid.UUID) (*postgres.MatchGoal, error)
	UpdateMatchGoalSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchGoal, error)
	DeleteMatchGoal(ctx context.Context, id uuid.UUID) error
}

type PostgresRepository interface {
	TeamRepository
	PlayerRepository
	MatchRepository
	MatchResultRepository
	MatchGoalRepository
}

func NewPostgres(db *gorm.DB) PostgresRepository {
	return internal.NewPostgresDB(db)
}
