package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/invoice"
	"github.com/turao/go-ddd/events"
)

type AddTaskToInvoiceCommandHandler struct {
	repository invoice.Repository
	eventStore events.EventStore
}

var _ application.AddTaskToInvoiceCommandHandler = (*AddTaskToInvoiceCommandHandler)(nil)

func NewAddTaskToInvoiceCommandHandler(repository invoice.Repository, es events.EventStore) *AddTaskToInvoiceCommandHandler {
	return &AddTaskToInvoiceCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AddTaskToInvoiceCommandHandler) Handle(ctx context.Context, req application.AddTaskToInvoiceCommand) error {
	i, err := h.repository.FindByID(ctx, req.InvoiceID)
	if err != nil {
		return err
	}

	ia, err := invoice.NewInvoiceAggregate(i, h.eventStore)
	if err != nil {
		return nil
	}

	err = ia.AddTask(req.TaskID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ia.Invoice)
	if err != nil {
		return err
	}

	return nil
}
