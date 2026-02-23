package internal

import (
	"ayo/impl/postgres/postgres"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *PostgresDB) CreatTeam(ctx context.Context, teamModel *postgres.Team) (*postgres.Team, error) {

	err := p.Db.WithContext(ctx).Create(teamModel).Error
	if err != nil {
		return nil, err
	}

	return teamModel, nil
}

func (p *PostgresDB) GetTeamByID(ctx context.Context, teamID uuid.UUID) (*postgres.Team, error) {
	var team postgres.Team
	if err := p.Db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", teamID).First(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &team, nil
}

func (p *PostgresDB) UpdateTeamSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Team, error) {
	var existing postgres.Team

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

	var updated postgres.Team
	if err := p.Db.WithContext(ctx).Where("id = ?", id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *PostgresDB) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	return p.Db.WithContext(ctx).
		Model(&postgres.Team{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
