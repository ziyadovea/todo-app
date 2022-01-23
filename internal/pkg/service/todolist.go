package service

import (
	"errors"
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/models"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (l *TodoListService) CreateNewList(userID int, list *models.TodoList) (int, error) {
	return l.repo.CreateNewList(userID, list)
}

func (l *TodoListService) GetAllLists(userID int) ([]*models.TodoList, error) {
	return l.repo.GetAllLists(userID)
}

func (l *TodoListService) DeleteAllLists(userID int) (int, error) {
	return l.repo.DeleteAllLists(userID)
}

func (l *TodoListService) GetListByID(userID, listID int) (*models.TodoList, error) {
	return l.repo.GetListByID(userID, listID)
}

func (l *TodoListService) DeleteListByID(userID, listID int) (int, error) {
	return l.repo.DeleteListByID(userID, listID)
}

func (l *TodoListService) UpdateListByID(userID, listID int, newList *models.NewTodoList) error {
	if newList.Title == nil && newList.Description == nil {
		return errors.New("update structure has no values")
	}
	return l.repo.UpdateListByID(userID, listID, newList)
}
