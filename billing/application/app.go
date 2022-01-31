package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateAccountCommand      CreateAccountCommandHandler
	AddTaskToUserCommand      AddTaskToUserCommandHandler
	RemoveTaskFromUserCommand RemoveTaskFromUserCommandHandler
}

type Queries struct {
}

type CreateAccountCommand struct {
	UserID string `json:"userId"`
}

type CreateAccountCommandHandler interface {
	Handle(ctx context.Context, req CreateAccountCommand) error
}

type AddTaskToUserCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type AddTaskToUserCommandHandler interface {
	Handle(ctx context.Context, req AddTaskToUserCommand) error
}

type RemoveTaskFromUserCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type RemoveTaskFromUserCommandHandler interface {
	Handle(ctx context.Context, req RemoveTaskFromUserCommand) error
}
