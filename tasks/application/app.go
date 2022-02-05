package application

import (
	"context"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTaskCommand        CreateTaskCommandHandler
	AssignToUserCommand      AssignToUserCommandHandler
	UnassignUserCommand      UnassignUserCommandHandler
	UpdateTitleCommand       UpdateTitleCommandHandler
	UpdateDescriptionCommand UpdateDescriptionCommandHandler
	UpdateStatusCommand      UpdateStatusCommandHandler
}

type Queries struct {
	TasksByProjectQuery      TasksByProjectQueryHandler
	TasksByAssignedUserQuery TasksByAssignedUserQueryHandler
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

type UnassignUserCommand struct {
	TaskID string `json:"taskId"`
}

type UnassignUserCommandHandler interface {
	Handle(ctx context.Context, req UnassignUserCommand) error
}

type UpdateTitleCommand struct {
	TaskID string `json:"taskId"`
	Title  string `json:"title"`
}

type UpdateTitleCommandHandler interface {
	Handle(ctx context.Context, req UpdateTitleCommand) error
}

type UpdateDescriptionCommand struct {
	TaskID      string `json:"taskId"`
	Description string `json:"description"`
}

type UpdateDescriptionCommandHandler interface {
	Handle(ctx context.Context, req UpdateDescriptionCommand) error
}

type UpdateStatusCommand struct {
	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

type UpdateStatusCommandHandler interface {
	Handle(ctx context.Context, req UpdateStatusCommand) error
}

// --- Queries ---
type Task struct {
	TaskID     string `json:"taskId"`
	AssignedTo string `json:"assignedTo"`
	Status     string `json:"status"`
}

type TasksByProjectQuery struct {
	ProjectID string `json:"projectId"`
}

type TasksByProjectResponse struct {
	ProjectID string `json:"projectId"`
	Tasks     []Task `json:"tasks"`
}

type TasksByProjectQueryHandler interface {
	Handle(ctx context.Context, req TasksByProjectQuery) (*TasksByProjectResponse, error)
}

type TasksByAssignedUserQuery struct {
	UserID string `json:"userId"`
}

type TasksByAssignedUserResponse struct {
	UserID string `json:"userId"`
	Tasks  []Task `json:"tasks"`
}

type TasksByAssignedUserQueryHandler interface {
	Handle(ctx context.Context, req TasksByAssignedUserQuery) (*TasksByAssignedUserResponse, error)
}
