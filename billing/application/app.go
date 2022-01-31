package application

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateInvoiceCommand         CreateInvoiceCommandHandler
	AddTaskToInvoiceCommand      AddTaskToInvoiceCommandHandler
	RemoveTaskFromInvoiceCommand RemoveTaskFromInvoiceCommandHandler
}

type Queries struct {
}

type CreateInvoiceCommand struct {
	UserID string `json:"userId"`
}

type CreateInvoiceCommandHandler interface {
	Handle(ctx context.Context, req CreateInvoiceCommand) error
}

type AddTaskToInvoiceCommand struct {
	InvoiceID string `json:"invoiceID"`
	TaskID    string `json:"taskId"`
}

type AddTaskToInvoiceCommandHandler interface {
	Handle(ctx context.Context, req AddTaskToInvoiceCommand) error
}

type RemoveTaskFromInvoiceCommand struct {
	InvoiceID string `json:"invoiceID"`
	TaskID    string `json:"taskId"`
}

type RemoveTaskFromInvoiceCommandHandler interface {
	Handle(ctx context.Context, req RemoveTaskFromInvoiceCommand) error
}
