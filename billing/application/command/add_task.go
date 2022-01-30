package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/invoice"
	"github.com/turao/go-ddd/events"
)

type AddTaskCommandHandler struct {
	repository invoice.Repository
	eventStore events.EventStore
}

var _ application.AddTaskCommandHandler = (*AddTaskCommandHandler)(nil)

func NewAddTaskCommandHandler(repository invoice.Repository, es events.EventStore) *AddTaskCommandHandler {
	return &AddTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AddTaskCommandHandler) Handle(ctx context.Context, req application.AddTaskCommand) error {
	i, err := h.repository.FindByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}

	ia, err := invoice.NewInvoiceAggregate(i, h.eventStore)
	if err != nil {
		return nil
	}

	err = ia.AddTask(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ia.Invoice)
	if err != nil {
		return err
	}

	return nil
}
