package internal

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"context"
	"time"
)

func (mmr *MatchResultManagement) CreatMatchResult(ctx context.Context, matchResultData *entity.MatchResult) (*entity.MatchResult, error) {

	matchModel := &postgres.MatchResult{
		MatchID:    matchResultData.MatchID,
		HomeScore:  matchResultData.HomeScore,
		AwayScore:  matchResultData.AwayScore,
		ReportedAt: time.Now(),
	}

	respCreateMatch, err := mmr.MatchResultRepo.CreatMatchResult(ctx, matchModel)
	if err != nil {
		return nil, err
	}

	resp := &entity.MatchResult{
		MatchID:    respCreateMatch.MatchID,
		HomeScore:  respCreateMatch.HomeScore,
		AwayScore:  respCreateMatch.AwayScore,
		ReportedAt: respCreateMatch.ReportedAt,
	}

	return resp, nil
}
