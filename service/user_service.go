package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/rzldimam28/simple-notes/models"
)

var Users = []models.User{}

type UserService struct {	
}

func (u *UserService) Create(newUser models.User) (models.User, error){
	id := uuid.NewString()

	user := models.User{
		Id:        id,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Users = append(Users, user)

	return user, nil
}

func (u *UserService) Lists() []models.User{
	return Users
}