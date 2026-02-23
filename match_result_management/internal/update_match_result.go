package internal

import (
	"ayo/entity"
	"context"
	"time"

	"github.com/google/uuid"
)

func (mrm *MatchResultManagement) UpdateMatchResult(ctx context.Context, id string, matchResultData *entity.MatchResult) (*entity.MatchResult, error) {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updatedToMatch := map[string]interface{}{
		"home_score":  matchResultData.HomeScore,
		"away_score":  matchResultData.AwayScore,
		"reported_at": time.Now,
	}

	respCreatMatch, err := mrm.MatchResultRepo.UpdateMatchResultSelective(ctx, matchID, updatedToMatch)
	if err != nil {
		return nil, err
	}

	resp := &entity.MatchResult{
		MatchID:    respCreatMatch.MatchID,
		HomeScore:  respCreatMatch.HomeScore,
		AwayScore:  respCreatMatch.AwayScore,
		ReportedAt: respCreatMatch.ReportedAt,
	}

	return resp, nil
}
