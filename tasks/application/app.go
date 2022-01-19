package application

import (
	"context"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTaskCommand   CreateTaskCommandHandler
	AssignToUserCommand AssignToUserCommandHandler
}

type Queries struct {
	TasksByProjectQuery TasksByProjectQueryHandler
}

// -- Commands --
type CreateTaskCommand struct {
	ProjectID   string `json:"projectId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskCommandHandler interface {
	Handle(ctx context.Context, req CreateTaskCommand) error
}

type AssignToUserCommand struct {
	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

type AssignToUserCommandHandler interface {
	Handle(ctx context.Context, req AssignToUserCommand) error
}

// --- Queries ---
type TasksByProjectQuery struct {
	ProjectID string `json:"projectId"`
}

type TasksByProjectResponse struct {
	ProjectID string `json:"projectId"`
	Tasks     []Task `json:"tasks"`
}

type Task struct {
	TaskID string `json:"taskId"`
}

type TasksByProjectQueryHandler interface {
	Handle(ctx context.Context, req TasksByProjectQuery) (*TasksByProjectResponse, error)
}

// -- Queriess --
