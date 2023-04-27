package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type PublisherRepository interface {
	FindAll(ctx context.Context) ([]model.Publisher, error)
	FindByIds(ctx context.Context, ids []uint64) ([]model.Publisher, error)
}
