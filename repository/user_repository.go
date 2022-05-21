package repository

import (
	"context"

	"github.com/rzldimam28/simple-notes/models/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) entity.User
	ListAll(ctx context.Context) []entity.User
	GetById(ctx context.Context, id int) entity.User
}