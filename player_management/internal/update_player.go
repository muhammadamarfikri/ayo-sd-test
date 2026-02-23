package internal

import (
	"ayo/entity"
	"context"

	"github.com/google/uuid"
)

func (pm *PlayerManagement) UpdatePlayer(ctx context.Context, id string, playerData *entity.Player) (*entity.Player, error) {

	teamID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updatedToTeam := map[string]interface{}{
		"name":         playerData.Name,
		"team_id":      playerData.TeamID,
		"height_cm":    playerData.HeightCM,
		"weight_kg":    playerData.WeightKG,
		"position":     playerData.Position,
		"shirt_number": playerData.ShirtNumber,
	}

	respCreatTeam, err := pm.PlayerRepo.UpdatePlayerSelective(ctx, teamID, updatedToTeam)
	if err != nil {
		return nil, err
	}

	return (*entity.Player)(respCreatTeam), nil
}
