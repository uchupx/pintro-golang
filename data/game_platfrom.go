package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type GamePlatformQuery struct {
	PerPage uint64
	Page    uint64
}

type GamePlatformRepository interface {
	// FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
	FindByPlatformIds(ctx context.Context, ids []uint64) ([]model.GamePlatform, error)
	FindByPublisherIds(ctx context.Context, ids []uint64) ([]model.GamePlatform, error)
}
