package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type GameQuery struct {
	PerPage uint64
	Page    uint64
}

type GameRepository interface {
	FindByQuery(ctx context.Context, query GameQuery) (*Collection, error)
	FindByIds(ctx context.Context, ids []uint64) ([]model.Game, error)
	Delete(ctx context.Context, game model.Game) (*int64, error)
	Update(ctx context.Context, game model.Game) (*int64, error)
	Insert(ctx context.Context, game model.Game) (*int64, error)
}
