package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/invoice"
	"github.com/turao/go-ddd/events"
)

type RemoveTaskCommandHandler struct {
	repository invoice.Repository
	eventStore events.EventStore
}

var _ application.RemoveTaskCommandHandler = (*RemoveTaskCommandHandler)(nil)

func NewRemoveTaskCommandHandler(repository invoice.Repository, es events.EventStore) *RemoveTaskCommandHandler {
	return &RemoveTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h RemoveTaskCommandHandler) Handle(ctx context.Context, req application.RemoveTaskCommand) error {
	i, err := h.repository.FindByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}

	ia, err := invoice.NewInvoiceAggregate(i, h.eventStore)
	if err != nil {
		return nil
	}

	err = ia.RemoveTask(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ia.Invoice)
	if err != nil {
		return err
	}

	return nil
}
