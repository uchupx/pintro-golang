package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type GamePublisherQuery struct {
	PerPage uint64
	Page    uint64
}

type GamePublisherRepository interface {
	FindByGameIds(ctx context.Context, ids []uint64) ([]model.GamePublisher, error)
	FindByPublisherIds(ctx context.Context, ids []uint64) ([]model.GamePublisher, error)
}
