package internal

import (
	"context"

	"github.com/google/uuid"
)

func (mm *MatchManagement) DeleteMatch(ctx context.Context, id string) error {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = mm.MatchRepo.DeleteMatch(ctx, matchID)
	if err != nil {
		return err
	}

	return nil
}
