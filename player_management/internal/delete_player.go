package internal

import (
	"context"

	"github.com/google/uuid"
)

func (pm *PlayerManagement) DeletePlayer(ctx context.Context, id string) error {

	playerID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = pm.PlayerRepo.DeletePlayer(ctx, playerID)
	if err != nil {
		return err
	}

	return nil
}
