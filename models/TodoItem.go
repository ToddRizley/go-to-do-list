package models

type TodoItem struct {
	ID          int    `json:"id"`
	TextContent string `json:"text_content"`
	Completed   bool   `json:"completed"`
}

func (b *TodoItem) TableName() string {
	return "todolist"
}
