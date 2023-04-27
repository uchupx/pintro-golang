package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type RegionRepository interface {
	// FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
	FindAll(ctx context.Context) ([]model.Region, error)
	FindByIds(ctx context.Context, ids []uint64) ([]model.Region, error)
}
