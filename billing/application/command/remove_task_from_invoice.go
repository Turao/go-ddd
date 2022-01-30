package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/invoice"
	"github.com/turao/go-ddd/events"
)

type RemoveTaskFromInvoiceCommandHandler struct {
	repository invoice.Repository
	eventStore events.EventStore
}

var _ application.RemoveTaskFromInvoiceCommandHandler = (*RemoveTaskFromInvoiceCommandHandler)(nil)

func NewRemoveTaskFromInvoiceCommandHandler(repository invoice.Repository, es events.EventStore) *RemoveTaskFromInvoiceCommandHandler {
	return &RemoveTaskFromInvoiceCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h RemoveTaskFromInvoiceCommandHandler) Handle(ctx context.Context, req application.RemoveTaskFromInvoiceCommand) error {
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
