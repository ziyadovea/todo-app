package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
)

type Authorization interface {
	CreateUser(*models.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      nil,
		TodoItem:      nil,
	}
}
