package internal

import (
	"ayo/impl/postgres/postgres"
	"context"

	"github.com/google/uuid"
)

type MatchRepository interface {
	CreatMatch(ctx context.Context, matchModel *postgres.Matches) (*postgres.Matches, error)
	GetMatchByID(ctx context.Context, matchID uuid.UUID) (*postgres.Matches, error)
	UpdateMatchSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Matches, error)
	DeleteMatch(ctx context.Context, id uuid.UUID) error
}

type MatchManagement struct {
	MatchRepo MatchRepository
}

func NewMatchManagement(mr MatchRepository) *MatchManagement {
	return &MatchManagement{mr}
}
