package internal

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"context"
)

func (tm *TeamManagement) CreatTeam(ctx context.Context, teamData *entity.Team) (*entity.Team, error) {

	teamModel := &postgres.Team{
		Name:                teamData.Name,
		LogoURL:             teamData.LogoURL,
		FoundedYear:         teamData.FoundedYear,
		HeadquartersAddress: teamData.HeadquartersAddress,
		HeadquartersCity:    teamData.HeadquartersCity,
	}

	respCreatTeam, err := tm.TeamRepo.CreatTeam(ctx, teamModel)
	if err != nil {
		return nil, err
	}

	return (*entity.Team)(respCreatTeam), nil
}
