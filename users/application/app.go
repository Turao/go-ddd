package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
}

// --- Commands ---
type RegisterUserCommand struct {
	Username string `json:"username"`
}

type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, req RegisterUserCommand) error
}
