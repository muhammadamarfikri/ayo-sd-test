package internal

import (
	"ayo/impl/postgres/postgres"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *PostgresDB) CreatPlayer(ctx context.Context, playerModel *postgres.Player) (*postgres.Player, error) {

	err := p.Db.WithContext(ctx).Create(playerModel).Error
	if err != nil {
		return nil, err
	}

	return playerModel, nil
}

func (p *PostgresDB) GetPlayerByID(ctx context.Context, playerID uuid.UUID) (*postgres.Player, error) {
	var player postgres.Player
	if err := p.Db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", playerID).First(&player).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (p *PostgresDB) UpdatePlayerSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Player, error) {
	var existing postgres.Player

	if err := p.Db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).Order("created_at DESC").First(&existing).Error; err != nil {
		return nil, err
	}

	updateData := make(map[string]interface{})

	for key, value := range updates {
		if value != nil {
			updateData[key] = value
		}
	}

	updateData["updated_at"] = time.Now()

	if err := p.Db.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated postgres.Player
	if err := p.Db.WithContext(ctx).Where("id = ?", id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *PostgresDB) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	return p.Db.WithContext(ctx).
		Model(&postgres.Player{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
