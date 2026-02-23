package playermanagement

import (
	"ayo/entity"
	"ayo/impl/postgres/postgres"
	"ayo/player_management/internal"
	"context"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	CreatPlayer(ctx context.Context, playerModel *postgres.Player) (*postgres.Player, error)
	GetPlayerByID(ctx context.Context, playerID uuid.UUID) (*postgres.Player, error)
	UpdatePlayerSelective(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*postgres.Player, error)
	DeletePlayer(ctx context.Context, id uuid.UUID) error
}

type PlayerManagement interface {
	CreatPlayer(ctx context.Context, playerData *entity.Player) (*entity.Player, error)
	GetPlayer(ctx context.Context, id string) (*entity.Player, error)
	UpdatePlayer(ctx context.Context, id string, playerData *entity.Player) (*entity.Player, error)
	DeletePlayer(ctx context.Context, id string) error
}

func NewPlayerManagement(pr PlayerRepository) PlayerManagement {
	return internal.NewPlayerManagement(pr)
}
