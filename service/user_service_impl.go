package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/repository"
)

type UserServiceImpl struct {	
	UserRepository repository.UserRepository
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate: validate,
	}
}

func (userService *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	user := entity.User{
		Username: request.Username,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUser := userService.UserRepository.Save(ctx, user)

	return helper.ToUserResponse(newUser)
}

func (userService *UserServiceImpl) Lists(ctx context.Context) []web.UserResponse {
	users := userService.UserRepository.ListAll(ctx)	
	return helper.ToUserResponses(users)
}