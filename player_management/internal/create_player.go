package internal

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"context"
)

func (pm *PlayerManagement) CreatPlayer(ctx context.Context, playerData *entity.Player) (*entity.Player, error) {

	playerModel := &postgres.Player{
		Name:        playerData.Name,
		TeamID:      playerData.TeamID,
		HeightCM:    playerData.HeightCM,
		WeightKG:    playerData.WeightKG,
		Position:    playerData.Position,
		ShirtNumber: playerData.ShirtNumber,
	}

	respCreatePlayer, err := pm.PlayerRepo.CreatPlayer(ctx, playerModel)
	if err != nil {
		return nil, err
	}

	return (*entity.Player)(respCreatePlayer), nil
}
