package internal

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"context"
)

func (mm *MatchManagement) CreatMatch(ctx context.Context, matchData *entity.Match) (*entity.Match, error) {

	matchModel := &postgres.Matches{
		MatchAt:    matchData.MatchAt,
		HomeTeamID: matchData.HomeTeamID,
		AwayTeamID: matchData.AwayTeamID,
		Status:     matchData.Status,
	}

	respCreateMatch, err := mm.MatchRepo.CreatMatch(ctx, matchModel)
	if err != nil {
		return nil, err
	}

	return (*entity.Match)(respCreateMatch), nil
}
