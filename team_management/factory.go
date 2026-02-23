package teammanagement

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"ayo/team_management/internal"
	"context"

	"github.com/google/uuid"
)

type TeamRepository interface {
	CreatTeam(ctx context.Context, teamModel *postgres.Team) (*postgres.Team, error)
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (*postgres.Team, error)
	UpdateTeamSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Team, error)
	DeleteTeam(ctx context.Context, id uuid.UUID) error
}

type TeamManagement interface {
	CreatTeam(ctx context.Context, teamData *entity.Team) (*entity.Team, error)
	GetTeam(ctx context.Context, id string) (*entity.Team, error)
	UpdateTeam(ctx context.Context, id string, teamData *entity.Team) (*entity.Team, error)
	DeleteTeam(ctx context.Context, id string) error
}

func NewTeamManagement(tr TeamRepository) TeamManagement {
	return internal.NewTeamManagement(tr)
}
