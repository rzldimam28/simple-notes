package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (userRepo *UserRepositoryImpl) Save(ctx context.Context, user entity.User) entity.User {
	SQL := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := userRepo.DB.ExecContext(ctx, SQL, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (userRepo *UserRepositoryImpl) ListAll(ctx context.Context) []entity.User {
	SQL := "SELECT id, username, password, created_at, updated_at FROM users"
	rows, err := userRepo.DB.QueryContext(ctx, SQL)
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

func (userRepo *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	SQL := "SELECT id, username, password, created_at, updated_at FROM users where username = ?"
	rows, err := userRepo.DB.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	var user entity.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	}
	return user, errors.New("can not find user")
}