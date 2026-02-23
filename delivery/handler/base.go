package handler

import (
	"ayo/entity"
	"context"
)

type TeamManagement interface {
	CreatTeam(ctx context.Context, teamData *entity.Team) (*entity.Team, error)
	GetTeam(ctx context.Context, id string) (*entity.Team, error)
	UpdateTeam(ctx context.Context, id string, teamData *entity.Team) (*entity.Team, error)
	DeleteTeam(ctx context.Context, id string) error
}

type PlayerManagement interface {
	CreatPlayer(ctx context.Context, playerData *entity.Player) (*entity.Player, error)
	GetPlayer(ctx context.Context, id string) (*entity.Player, error)
	UpdatePlayer(ctx context.Context, id string, playerData *entity.Player) (*entity.Player, error)
	DeletePlayer(ctx context.Context, id string) error
}

type MatchManagement interface {
	CreatMatch(ctx context.Context, matchData *entity.Match) (*entity.Match, error)
	GetMatchByID(ctx context.Context, id string) (*entity.Match, error)
	UpdateMatch(ctx context.Context, id string, matchData *entity.Match) (*entity.Match, error)
	DeleteMatch(ctx context.Context, id string) error
}

type MatchResultManagement interface {
	CreatMatchResult(ctx context.Context, matchResultData *entity.MatchResult) (*entity.MatchResult, error)
	GetMatchResultByID(ctx context.Context, id string) (*entity.MatchResult, error)
	UpdateMatchResult(ctx context.Context, id string, matchResultData *entity.MatchResult) (*entity.MatchResult, error)
	DeleteMatchResult(ctx context.Context, id string) error
}

type Handler struct {
	TM  TeamManagement
	PM  PlayerManagement
	MM  MatchManagement
	MRM MatchResultManagement
}

func NewHandler(tm TeamManagement, pm PlayerManagement, mm MatchManagement, mrm MatchResultManagement) *Handler {
	return &Handler{
		TM:  tm,
		PM:  pm,
		MM:  mm,
		MRM: mrm,
	}
}
