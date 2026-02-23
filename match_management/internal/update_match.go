package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (mm *MatchManagement) UpdateMatch(ctx context.Context, id string, matchData *entity.Match) (*entity.Match, error) {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updatedToMatch := map[string]interface{}{
		"status": matchData.Status,
	}

	respCreatMatch, err := mm.MatchRepo.UpdateMatchSelective(ctx, matchID, updatedToMatch)
	if err != nil {
		return nil, err
	}

	return (*entity.Match)(respCreatMatch), nil
}
