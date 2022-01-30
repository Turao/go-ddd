package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/invoice"
	"github.com/turao/go-ddd/events"
)

type CreateInvoiceCommandHandler struct {
	repository invoice.Repository
	eventStore events.EventStore
}

var _ application.CreateInvoiceCommandHandler = (*CreateInvoiceCommandHandler)(nil)

func NewCreateInvoiceCommandHandler(repository invoice.Repository, es events.EventStore) *CreateInvoiceCommandHandler {
	return &CreateInvoiceCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h CreateInvoiceCommandHandler) Handle(ctx context.Context, req application.CreateInvoiceCommand) error {
	ia, err := invoice.NewInvoiceAggregate(nil, h.eventStore)
	if err != nil {
		return nil
	}

	err = ia.CreateInvoice(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ia.Invoice)
	if err != nil {
		return err
	}

	return nil
}
