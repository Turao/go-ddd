package application

import (
	"context"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}

type Queries struct{}

// -- Commands --
type CreateTaskCommand struct {
	ID          string `json:"taskId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskCommandHandler interface {
	Handle(ctx context.Context, req CreateTaskCommand) error
}

// -- Queriess --
