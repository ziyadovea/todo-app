package models

type TodoItem struct {
	ID          int    `json:"id"`
	ListID      int    `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
