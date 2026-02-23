package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (tm *TeamManagement) GetTeam(ctx context.Context, id string) (*entity.Team, error) {

	teamID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	respCreatTeam, err := tm.TeamRepo.GetTeamByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return (*entity.Team)(respCreatTeam), nil
}
