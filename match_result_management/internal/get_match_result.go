package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (mrm *MatchResultManagement) GetMatchResultByID(ctx context.Context, id string) (*entity.MatchResult, error) {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	respGetMatch, err := mrm.MatchResultRepo.GetMatchResultByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	resp := &entity.MatchResult{
		MatchID:    respGetMatch.MatchID,
		HomeScore:  respGetMatch.HomeScore,
		AwayScore:  respGetMatch.AwayScore,
		ReportedAt: respGetMatch.ReportedAt,
	}

	return resp, nil
}
