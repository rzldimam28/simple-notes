package service

import (
	"context"

	"github.com/rzldimam28/simple-notes/models/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Lists(ctx context.Context) []web.UserResponse
}