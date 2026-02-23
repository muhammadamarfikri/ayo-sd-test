package matchresultmanagement

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"ayo/match_result_management/internal"
	"context"

	"github.com/google/uuid"
)

type MatchResultRepository interface {
	CreatMatchResult(ctx context.Context, matchModel *postgres.MatchResult) (*postgres.MatchResult, error)
	GetMatchResultByID(ctx context.Context, matchID uuid.UUID) (*postgres.MatchResult, error)
	UpdateMatchResultSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchResult, error)
	DeleteMatchResult(ctx context.Context, id uuid.UUID) error
}

type MatchResultManagement interface {
	CreatMatchResult(ctx context.Context, matchResultData *entity.MatchResult) (*entity.MatchResult, error)
	GetMatchResultByID(ctx context.Context, id string) (*entity.MatchResult, error)
	UpdateMatchResult(ctx context.Context, id string, matchResultData *entity.MatchResult) (*entity.MatchResult, error)
	DeleteMatchResult(ctx context.Context, id string) error
}

func NewMatchResultManagement(mr MatchResultRepository) MatchResultManagement {
	return internal.NewMatchResultManagement(mr)
}
