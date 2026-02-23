package internal

import (
	"ayo/impl/postgres/postgres"
	"context"

	"github.com/google/uuid"
)

type TeamRepository interface {
	CreatTeam(ctx context.Context, teamModel *postgres.Team) (*postgres.Team, error)
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (*postgres.Team, error)
	UpdateTeamSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Team, error)
	DeleteTeam(ctx context.Context, id uuid.UUID) error
}

type TeamManagement struct {
	TeamRepo TeamRepository
}

func NewTeamManagement(tr TeamRepository) *TeamManagement {
	return &TeamManagement{tr}
}
