package matchmanagement

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"ayo/match_management/internal"
	"context"

	"github.com/google/uuid"
)

type MatchRepository interface {
	CreatMatch(ctx context.Context, matchModel *postgres.Matches) (*postgres.Matches, error)
	GetMatchByID(ctx context.Context, matchID uuid.UUID) (*postgres.Matches, error)
	UpdateMatchSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Matches, error)
	DeleteMatch(ctx context.Context, id uuid.UUID) error
}

type MatchManagement interface {
	CreatMatch(ctx context.Context, matchData *entity.Match) (*entity.Match, error)
	GetMatchByID(ctx context.Context, id string) (*entity.Match, error)
	UpdateMatch(ctx context.Context, id string, matchData *entity.Match) (*entity.Match, error)
	DeleteMatch(ctx context.Context, id string) error
}

func NewMatchManagement(mr MatchRepository) MatchManagement {
	return internal.NewMatchManagement(mr)
}
