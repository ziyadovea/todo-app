package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user *models.User) (int, error) {
	err := ap.db.QueryRow(
		`insert into users (Name, Username, Password_hash) values ($1, $2, $3) returning ID`,
		user.Name,
		user.Username,
		user.Password,
	).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
