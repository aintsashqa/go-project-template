package handler

import (
	"errors"
	"log"
	"time"

	"github.com/aintsashqa/go-project-template/types"
)

type TodoHandler struct {
	index int64
}

func NewTodoHandler() *TodoHandler {
	return new(TodoHandler)
}

func (th *TodoHandler) Create(args types.CreateTodoArgs, reply *types.Todo) error {
	log.Print("Executing TodoHandler.Create method")

	th.index += 1
	*reply = types.Todo{
		ID:          th.index,
		Title:       args.Title,
		Description: args.Description,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	return nil
}

func (th *TodoHandler) Complete(_ int64, _ *bool) error {
	log.Print("Executing TodoHandler.Complete method")
	return errors.New("todo is already completed")
}
