package service

import (
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/models"
)

type Authorization interface {
	CreateUser(*models.User) (int, error)
}

type TodoList interface {
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
		TodoList:      nil,
		TodoItem:      nil,
	}
}
