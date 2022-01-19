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
	ListUsersQuery ListUsersQueryHandler
}

// --- Commands ---
type RegisterUserCommand struct {
	Username string `json:"username"`
}

type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, req RegisterUserCommand) error
}

// --- Queries ---
type ListUsersQuery struct{}

type ListUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

type ListUsersQueryHandler interface {
	Handle(ctx context.Context, req ListUsersQuery) (*ListUsersResponse, error)
}
