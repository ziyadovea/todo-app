package service

import (
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/models"
)

type Authorization interface {
	CreateUser(*models.User) (int, error)
	GenerateToken(string, string) (string, error)
	ParseToken(string) (int, error)
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

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      nil,
	}
}
