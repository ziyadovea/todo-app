package models

type TodoList struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
