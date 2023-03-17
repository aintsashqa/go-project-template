package types

import "time"

type Todo struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type CreateTodoArgs struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}
