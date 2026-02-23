package internal

import (
	"context"

	"github.com/google/uuid"
)

func (tm *TeamManagement) DeleteTeam(ctx context.Context, id string) error {

	teamID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = tm.TeamRepo.DeleteTeam(ctx, teamID)
	if err != nil {
		return err
	}

	return nil
}
