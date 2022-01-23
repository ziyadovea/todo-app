package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ziyadovea/todo-app/models"
	"strconv"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (l *TodoListPostgres) CreateNewList(userID int, list *models.TodoList) (int, error) {
	var id int
	err := l.db.QueryRowx(
		"insert into todo_lists (Title, Description, User_ID) values ($1, $2, $3) returning id",
		list.Title,
		list.Description,
		userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (l *TodoListPostgres) GetAllLists(userID int) ([]*models.TodoList, error) {
	rows, err := l.db.Queryx("select * from todo_lists where user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	lists := make([]*models.TodoList, 0)
	for rows.Next() {
		list := &models.TodoList{}
		err := rows.StructScan(list)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (l *TodoListPostgres) DeleteAllLists(userID int) (int, error) {
	res, err := l.db.Exec("delete from todo_lists")
	if err != nil {
		return 0, err
	}
	affectedCount, err := res.RowsAffected()
	return int(affectedCount), err
}

func (l *TodoListPostgres) GetListByID(userID, listID int) (*models.TodoList, error) {
	list := &models.TodoList{}

	err := l.db.QueryRowx(
		"select * from todo_lists where id=$1 and user_id=$2",
		listID,
		userID,
	).StructScan(list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (l *TodoListPostgres) DeleteListByID(userID int, listID int) (int, error) {
	res, err := l.db.Exec("delete from todo_lists where id=$1 and user_id=$2", listID, userID)
	if err != nil {
		return 0, err
	}
	affectedCount, err := res.RowsAffected()
	return int(affectedCount), err
}

func (l *TodoListPostgres) UpdateListByID(userID int, listID int, newList *models.NewTodoList) error {

	queryStr := &strings.Builder{}
	queryStr.WriteString("update todo_lists set ")
	args := make([]interface{}, 0)

	if newList.Title != nil {
		args = append(args, *newList.Title)
		queryStr.WriteString("title=$" + strconv.Itoa(len(args)) + " ")
	}

	if newList.Description != nil {
		args = append(args, *newList.Description)
		queryStr.WriteString(", description=$" + strconv.Itoa(len(args)) + " ")
	}

	args = append(args, listID)
	queryStr.WriteString("where id=$" + strconv.Itoa(len(args)) + " and ")

	args = append(args, userID)
	queryStr.WriteString("user_id=$" + strconv.Itoa(len(args)))

	_, err := l.db.Exec(
		queryStr.String(),
		args...,
	)

	if err != nil {
		return err
	}

	return nil
}
