package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type RegionSalesRepository interface {
	FindByGamePlatformIds(ctx context.Context, ids []uint64) ([]model.RegionSales, error)
}
