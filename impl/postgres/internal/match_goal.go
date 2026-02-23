package internal

import (
	"ayo/impl/postgres/postgres"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *PostgresDB) CreatMatchGoal(ctx context.Context, matchGoalModel *postgres.MatchGoal) (*postgres.MatchGoal, error) {

	err := p.Db.WithContext(ctx).Create(matchGoalModel).Error
	if err != nil {
		return nil, err
	}

	return matchGoalModel, nil
}

func (p *PostgresDB) GetMatchGoalByID(ctx context.Context, matchID uuid.UUID) (*postgres.MatchGoal, error) {
	var matchGoal postgres.MatchGoal
	if err := p.Db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", matchID).First(&matchGoal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &matchGoal, nil
}

func (p *PostgresDB) UpdateMatchGoalSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchGoal, error) {
	var existing postgres.MatchGoal

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

	var updated postgres.MatchGoal
	if err := p.Db.WithContext(ctx).Where("id = ?", id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *PostgresDB) DeleteMatchGoal(ctx context.Context, id uuid.UUID) error {
	return p.Db.WithContext(ctx).
		Model(&postgres.MatchGoal{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
