package repository

import (
	"context"

	"github.com/rzldimam28/simple-notes/models/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) entity.User
	ListAll(ctx context.Context) []entity.User
	GetByUsername(ctx context.Context, username string) (entity.User, error)
}