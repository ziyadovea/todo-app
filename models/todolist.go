package models

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

// NewTodoList - Структура для запроса на обновление
type NewTodoList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
