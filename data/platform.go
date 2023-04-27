package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type PlatformRepository interface {
	// FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
	FindAll(ctx context.Context) ([]model.Platform, error)
	FindByIds(ctx context.Context, ids []uint64) ([]model.Platform, error)
}
