package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

// type Genre struct {
// 	PerPage uint64
// 	Page    uint64
// }

type GenreRepository interface {
	// FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
	FindAll(ctx context.Context) ([]model.Genre, error)
	FindByIds(ctx context.Context, ids []uint64) ([]model.Genre, error)
}
