package service

import (
	"github.com/ziyadovea/todo-app/internal/app/repository"
	"github.com/ziyadovea/todo-app/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

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
	CreateNewItem(int, int, *models.TodoItem) (int, error)
	GetAllItems(int, int) ([]*models.TodoItem, error)
	DeleteAllItems(int, int) (int, error)
	GetItemByID(int, int, int) (*models.TodoItem, error)
	DeleteItemByID(int, int, int) (int, error)
	UpdateItemByID(int, int, int, *models.NewTodoItem) error
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
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
