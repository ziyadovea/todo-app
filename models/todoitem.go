package models

type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	ListID      int    `json:"list_id" db:"list_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type NewTodoItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}
