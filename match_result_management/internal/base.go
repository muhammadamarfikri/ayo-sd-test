package internal

import (
	"ayo/impl/postgres/postgres"
	"context"

	"github.com/google/uuid"
)

type MatchResultRepository interface {
	CreatMatchResult(ctx context.Context, matchModel *postgres.MatchResult) (*postgres.MatchResult, error)
	GetMatchResultByID(ctx context.Context, matchID uuid.UUID) (*postgres.MatchResult, error)
	UpdateMatchResultSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchResult, error)
	DeleteMatchResult(ctx context.Context, id uuid.UUID) error
}

type MatchResultManagement struct {
	MatchResultRepo MatchResultRepository
}

func NewMatchResultManagement(mrr MatchResultRepository) *MatchResultManagement {
	return &MatchResultManagement{mrr}
}
