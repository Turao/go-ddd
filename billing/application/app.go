package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterUserCommand RegisterUserCommandHandler
	AssignTaskCommand   AssignTaskCommandHandler
	UnassignTaskCommand UnassignTaskCommandHandler
}

type Queries struct {
}

type RegisterUserCommand struct {
	UserID string `json:"userId"`
}

type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, req RegisterUserCommand) error
}

type AssignTaskCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type AssignTaskCommandHandler interface {
	Handle(ctx context.Context, req AssignTaskCommand) error
}

type UnassignTaskCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type UnassignTaskCommandHandler interface {
	Handle(ctx context.Context, req UnassignTaskCommand) error
}
