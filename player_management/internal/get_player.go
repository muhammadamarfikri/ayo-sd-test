package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (pm *PlayerManagement) GetPlayer(ctx context.Context, id string) (*entity.Player, error) {

	teamID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	respGetPlayer, err := pm.PlayerRepo.GetPlayerByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return (*entity.Player)(respGetPlayer), nil
}
