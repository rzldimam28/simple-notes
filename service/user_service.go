package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Validate *validator.Validate
}

func (userService *UserService) Create(request web.UserCreateRequest) web.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	user := entity.User{
		Username: request.Username,
		Password: request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUser := userService.UserRepository.Save(user)
	return helper.ToUserResponse(newUser)
}

func (userService *UserService) Get(username string) web.UserResponse {
	user, err := userService.UserRepository.Get(username)
	helper.PanicIfError(err)
	return helper.ToUserResponse(user)
}

func (userService *UserService) List() []web.UserResponse {
	users := userService.UserRepository.List()
	return helper.ToUserResponses(users)
}