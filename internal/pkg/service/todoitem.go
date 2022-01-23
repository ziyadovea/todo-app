package service

import (
	"errors"
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/models"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (i *TodoItemService) CreateNewItem(userID int, listID int, item *models.TodoItem) (int, error) {
	// Сначала надо проверить, принадлежит ли данный список данному пользователю
	if !i.isUserList(userID, listID) {
		return 0, errors.New("user does not have such a list")
	}
	return i.repo.CreateNewItem(listID, item)
}

func (i *TodoItemService) GetAllItems(userID int, listID int) ([]*models.TodoItem, error) {
	return i.repo.GetAllItems(userID, listID)
}

func (i *TodoItemService) DeleteAllItems(userID int, listID int) (int, error) {
	// Сначала надо проверить, принадлежит ли данный список данному пользователю
	if !i.isUserList(userID, listID) {
		return 0, errors.New("user does not have such a list")
	}
	return i.repo.DeleteAllItems(listID)
}

func (i *TodoItemService) GetItemByID(userID int, listID int, itemID int) (*models.TodoItem, error) {
	return i.repo.GetItemByID(userID, listID, itemID)
}

func (i *TodoItemService) DeleteItemByID(userID int, listID int, itemID int) (int, error) {
	// Сначала надо проверить, принадлежит ли данный список данному пользователю
	if !i.isUserList(userID, listID) {
		return 0, errors.New("user does not have such a list")
	}
	return i.repo.DeleteItemByID(listID, itemID)
}

func (i *TodoItemService) UpdateItemByID(userID int, listID int, itemID int, newItem *models.NewTodoItem) error {
	// Сначала надо проверить, принадлежит ли данный список данному пользователю
	if !i.isUserList(userID, listID) {
		return errors.New("user does not have such a list")
	}
	if newItem.Title == nil && newItem.Description == nil && newItem.Done == nil {
		return errors.New("update structure has no values")
	}
	return i.repo.UpdateItemByID(listID, itemID, newItem)
}

func (i *TodoItemService) isUserList(userID, listID int) bool {
	_, err := i.listRepo.GetListByID(userID, listID)
	return err == nil
}
