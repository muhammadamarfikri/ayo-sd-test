package internal

import (
	"context"

	"github.com/google/uuid"
)

func (mrm *MatchResultManagement) DeleteMatchResult(ctx context.Context, id string) error {

	matchID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = mrm.MatchResultRepo.DeleteMatchResult(ctx, matchID)
	if err != nil {
		return err
	}

	return nil
}
