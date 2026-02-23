package internal

import (
	"ayo/entity"
	"context"
	"time"

	"github.com/google/uuid"
)

func (tm *TeamManagement) UpdateTeam(ctx context.Context, id string, teamData *entity.Team) (*entity.Team, error) {

	teamID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updatedToTeam := map[string]interface{}{
		"name":                 teamData.Name,
		"logo_url":             teamData.LogoURL,
		"founded_year":         teamData.FoundedYear,
		"headquarters_address": teamData.HeadquartersAddress,
		"headquarters_city":    teamData.HeadquartersCity,
		"updated_at":           time.Now(),
	}

	respCreatTeam, err := tm.TeamRepo.UpdateTeamSelective(ctx, teamID, updatedToTeam)
	if err != nil {
		return nil, err
	}

	return (*entity.Team)(respCreatTeam), nil
}
