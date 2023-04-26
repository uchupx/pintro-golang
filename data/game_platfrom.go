package data

import "context"

type GamePlatformQuery struct {
	PerPage uint64
	Page    uint64
}

type GamePlatformRepository interface {
	FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
}
