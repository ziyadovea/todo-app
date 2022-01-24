package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
	"strconv"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func (i *TodoItemPostgres) CreateNewItem(listID int, item *models.TodoItem) (int, error) {
	err := i.db.QueryRowx(
		"insert into todo_items (title, description, done, list_id) values ($1, $2, $3, $4) returning id",
		item.Title,
		item.Description,
		item.Done,
		listID,
	).Scan(&item.ID)

	if err != nil {
		return 0, err
	}

	return item.ID, nil
}

func (i *TodoItemPostgres) GetAllItems(userID, listID int) ([]*models.TodoItem, error) {
	rows, err := i.db.Queryx(
		"select todo_items.id, todo_items.title, todo_items.Description, todo_items.done, todo_items.list_id "+
			"from todo_items inner join todo_lists on todo_items.list_id=todo_lists.id "+
			"where todo_items.list_id=$1 and todo_lists.user_id=$2",
		listID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	items := make([]*models.TodoItem, 0)
	for rows.Next() {
		item := &models.TodoItem{}
		if err = rows.StructScan(item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (i *TodoItemPostgres) DeleteAllItems(listID int) (int, error) {
	res, err := i.db.Exec("delete from todo_items where list_id=$1", listID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (i *TodoItemPostgres) GetItemByID(userID int, listID int, itemID int) (*models.TodoItem, error) {
	item := &models.TodoItem{}
	err := i.db.QueryRowx(
		"select todo_items.id, todo_items.title, todo_items.Description, todo_items.done, todo_items.list_id "+
			"from todo_items "+
			"inner join todo_lists "+
			"on todo_items.list_id=todo_lists.id "+
			"where todo_items.id=$1 and "+
			"todo_lists.id=$2 and "+
			"todo_lists.user_id=$3",
		itemID,
		listID,
		userID,
	).StructScan(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *TodoItemPostgres) DeleteItemByID(listID int, itemID int) (int, error) {
	res, err := i.db.Exec(
		"delete from todo_items "+
			"where id=$1 and "+
			"list_id=$2",
		itemID,
		listID,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (i *TodoItemPostgres) UpdateItemByID(listID int, itemID int, newItem *models.NewTodoItem) error {
	queryStr := &strings.Builder{}
	queryStr.WriteString("update todo_items set ")
	args := make([]interface{}, 0)
	setValues := make([]string, 0)

	if newItem.Title != nil {
		args = append(args, *newItem.Title)
		setValues = append(setValues, "title=$"+strconv.Itoa(len(args)))
	}

	if newItem.Description != nil {
		args = append(args, *newItem.Description)
		setValues = append(setValues, "description=$"+strconv.Itoa(len(args)))
	}

	if newItem.Done != nil {
		args = append(args, *newItem.Done)
		setValues = append(setValues, "done=$"+strconv.Itoa(len(args)))
	}

	queryStr.WriteString(strings.Join(setValues, ", "))

	args = append(args, itemID)
	queryStr.WriteString("where id=$" + strconv.Itoa(len(args)) + " and ")

	args = append(args, listID)
	queryStr.WriteString("list_id=$" + strconv.Itoa(len(args)))

	fmt.Println(queryStr)

	_, err := i.db.Exec(
		queryStr.String(),
		args...,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}
