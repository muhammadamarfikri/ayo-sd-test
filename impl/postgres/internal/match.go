package internal

import (
	"ayo/impl/postgres/postgres"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *PostgresDB) CreatMatch(ctx context.Context, matchModel *postgres.Matches) (*postgres.Matches, error) {

	err := p.Db.WithContext(ctx).Create(matchModel).Error
	if err != nil {
		return nil, err
	}

	return matchModel, nil
}

func (p *PostgresDB) GetMatchByID(ctx context.Context, matchID uuid.UUID) (*postgres.Matches, error) {
	var match postgres.Matches
	if err := p.Db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", matchID).First(&match).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &match, nil
}

func (p *PostgresDB) UpdateMatchSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Matches, error) {
	var existing postgres.Matches

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

	var updated postgres.Matches
	if err := p.Db.WithContext(ctx).Where("id = ?", id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *PostgresDB) DeleteMatch(ctx context.Context, id uuid.UUID) error {
	return p.Db.WithContext(ctx).
		Model(&postgres.Matches{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
