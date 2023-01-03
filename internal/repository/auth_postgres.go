package repository

import (
	"database/sql"
	"fmt"

	"todo-app/internal/model"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a *AuthPostgres) CreateUser(user model.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (fullname, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	var id int
	err := a.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(id), nil
}

func (a *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, fullname, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2", userTable)

	err := a.db.QueryRow(query, username, password).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}
	return user, nil
}
