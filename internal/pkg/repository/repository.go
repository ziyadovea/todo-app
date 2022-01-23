package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
)

var (
	ErrorIncorrectUsernameOrPassword = errors.New("incorrect username or password")
)

type Authorization interface {
	CreateUser(*models.User) (int, error)
	GetUser(username, password string) (*models.User, error)
}

type TodoList interface {
	CreateNewList(int, *models.TodoList) (int, error)
	GetAllLists(int) ([]*models.TodoList, error)
	DeleteAllLists(int) (int, error)
	GetListByID(int, int) (*models.TodoList, error)
	DeleteListByID(int, int) (int, error)
	UpdateListByID(int, int, *models.NewTodoList) error
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
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      nil,
	}
}
