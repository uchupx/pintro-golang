package data

import "context"

type GameQuery struct {
	PerPage uint64
	Page    uint64
}

type GameRepository interface {
	FindByQuery(ctx context.Context, query GameQuery) (*Collection, error)
}
