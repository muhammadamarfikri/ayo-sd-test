package internal

import (
	"ayo/impl/postgres/postgres"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *PostgresDB) CreatMatchResult(ctx context.Context, matchModel *postgres.MatchResult) (*postgres.MatchResult, error) {

	err := p.Db.WithContext(ctx).Create(matchModel).Error
	if err != nil {
		return nil, err
	}

	return matchModel, nil
}

func (p *PostgresDB) GetMatchResultByID(ctx context.Context, matchResultID uuid.UUID) (*postgres.MatchResult, error) {
	var matchResult postgres.MatchResult
	if err := p.Db.WithContext(ctx).Where("match_id = ? AND deleted_at IS NULL", matchResultID).First(&matchResult).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &matchResult, nil
}

func (p *PostgresDB) UpdateMatchResultSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.MatchResult, error) {
	var existing postgres.MatchResult

	if err := p.Db.WithContext(ctx).Where("match_id = ? AND deleted_at IS NULL", id).Order("reported_at DESC").First(&existing).Error; err != nil {
		return nil, err
	}

	updateData := make(map[string]interface{})

	for key, value := range updates {
		if value != nil {
			updateData[key] = value
		}
	}

	updateData["reported_at"] = time.Now()

	if err := p.Db.WithContext(ctx).Where("match_id = ? AND deleted_at IS NULL", id).Model(&existing).Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated postgres.MatchResult
	if err := p.Db.WithContext(ctx).Where("match_id = ?", id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *PostgresDB) DeleteMatchResult(ctx context.Context, id uuid.UUID) error {
	return p.Db.WithContext(ctx).
		Model(&postgres.MatchResult{}).
		Where("match_id = ?", id).
		Update("deleted_at", time.Now()).Error
}
