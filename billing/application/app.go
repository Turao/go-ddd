package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateInvoiceCommand CreateInvoiceCommandHandler
	AddTaskCommand       AddTaskCommandHandler
	RemoveTaskCommand    RemoveTaskCommandHandler
}

type Queries struct {
}

type CreateInvoiceCommand struct {
	UserID string `json:"userId"`
}

type CreateInvoiceCommandHandler interface {
	Handle(ctx context.Context, req CreateInvoiceCommand) error
}

type AddTaskCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type AddTaskCommandHandler interface {
	Handle(ctx context.Context, req AddTaskCommand) error
}

type RemoveTaskCommand struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
}

type RemoveTaskCommandHandler interface {
	Handle(ctx context.Context, req RemoveTaskCommand) error
}
