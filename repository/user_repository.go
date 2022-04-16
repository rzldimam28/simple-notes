package repository

import (
	"database/sql"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func (userRepo *UserRepository) Save(user entity.User) entity.User {
	SQL := "INSERT INTO users (first_name, last_name) VALUES (?, ?)"
	result, err := userRepo.DB.Exec(SQL, user.FirstName, user.LastName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (userRepo *UserRepository) ListAll() []entity.User {
	SQL := "SELECT id, first_name, last_name, created_at, updated_at FROM users"
	rows, err := userRepo.DB.Query(SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}