package repository

import (
	"database/sql"
	"errors"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func (userRepository *UserRepository) Save(user entity.User) entity.User {
	SQL := "INSERT INTO users(username, password, created_at, updated_at) VALUES(?, ?, ?, ?)"
	result, err := userRepository.DB.Exec(SQL, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	helper.PanicIfError(err)
	
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (userRepository *UserRepository) Get(id int) (entity.User, error) {
	SQL := "SELECT id, username, password, created_at, updated_at FROM users WHERE id = ?"
	rows, err := userRepository.DB.Query(SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var user entity.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	}
	return user, errors.New("Can not Find User")
}

func (userRepository *UserRepository) List() []entity.User {
	SQL := "SELECT id, username, password, created_at, updated_at FROM users"
	rows, err := userRepository.DB.Query(SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}