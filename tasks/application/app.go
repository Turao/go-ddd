package application

import (
	"context"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTaskCommand CreateTaskCommandHandler
}

type Queries struct {
	TasksByProjectQuery TasksByProjectQueryHandler
}

// -- Commands --
type CreateTaskCommand struct {
	ID          string `json:"taskId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskCommandHandler interface {
	Handle(ctx context.Context, req CreateTaskCommand) error
}
type TasksByProjectQuery struct {
	ProjectID string `json:"projectId"`
}

type TasksByProjectResponse struct {
	ProjectID string `json:"projectId"`
}

type TasksByProjectQueryHandler interface {
	Handle(ctx context.Context, req TasksByProjectQuery) (*TasksByProjectResponse, error)
}

// -- Queriess --
