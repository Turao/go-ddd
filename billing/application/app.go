package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterUserCommand RegisterUserCommandHandler
}

type Queries struct {
}

type RegisterUserCommand struct {
	UserID string `json:"userId"`
}

type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, req RegisterUserCommand) error
}
