package data

import "context"

type GamePublisherQuery struct {
	PerPage uint64
	Page    uint64
}

type GamePublisherRepository interface {
	FindQuery(ctx context.Context, query GamePlatformQuery) (*Collection, error)
}
