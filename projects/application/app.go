package application

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProject     CreateProjectHandler
	UpdateProject     UpdateProjectHandler
	DeleteProject     DeleteProjectHandler
	AddTaskCommand    AddTaskCommandHandler
	RemoveTaskCommand RemoveTaskCommandHandler
}

type Queries struct {
	FindProject FindProjectQueryHandler
}

// -- Commands --
type CreateProjectCommand struct {
	Name string `json:"name"`
}

type CreateProjectHandler interface {
	Handle(ctx context.Context, req CreateProjectCommand) error
}

type DeleteProjectCommand struct {
	ID project.ProjectID `json:"id"`
}

type DeleteProjectHandler interface {
	Handle(ctx context.Context, req DeleteProjectCommand) error
}

type UpdateProjectCommand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateProjectHandler interface {
	Handle(ctx context.Context, req UpdateProjectCommand) error
}

type AddTaskCommand struct {
	ID     string `json:"id"`
	TaskID string `json:"taskId"`
}

type AddTaskCommandHandler interface {
	Handle(ctx context.Context, req AddTaskCommand) error
}

type RemoveTaskCommand struct {
	ID     string `json:"id"`
	TaskID string `json:"taskId"`
}

type RemoveTaskCommandHandler interface {
	Handle(ctx context.Context, req RemoveTaskCommand) error
}

// -- Queries --

type FindProjectQuery struct {
	ID string `json:"id"`
}

type FindProjectResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type FindProjectQueryHandler interface {
	Handle(ctx context.Context, req FindProjectQuery) (*FindProjectResponse, error)
}
