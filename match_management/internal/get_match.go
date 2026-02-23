package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (mm *MatchManagement) GetMatchByID(ctx context.Context, id string) (*entity.Match, error) {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	respGetMatch, err := mm.MatchRepo.GetMatchByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return (*entity.Match)(respGetMatch), nil
}
