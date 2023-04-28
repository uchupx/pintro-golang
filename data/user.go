package data

import (
	"context"

	"github.com/uchupx/pintro-golang/data/model"
)

type UserRepoitory interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Insert(ctx context.Context, user model.User) (*int64, error)
	Update(ctx context.Context, user model.User) (*int64, error)
	Delete(ctx context.Context, user model.User) (*int64, error)
}
