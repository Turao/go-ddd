package invoice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/domain/user"
)

type InvoiceAggregate struct {
	Invoice *Invoice
	events  events.EventStore
}

func NewInvoiceAggregate(invoice *Invoice, es events.EventStore) (*InvoiceAggregate, error) {
	return &InvoiceAggregate{
		Invoice: invoice,
		events:  es,
	}, nil
}

func (ia *InvoiceAggregate) HandleEvent(event events.DomainEvent) error {
	switch e := event.(type) {
	case InvoiceCreatedEvent:
		i, err := NewInvoice(e.AggregateID(), e.UserID)
		if err != nil {
			return err
		}
		ia.Invoice = i
		return nil
	case TaskAddedEvent:
		return ia.AddTask(e.TaskID)
	case TaskRemovedEvent:
		return ia.RemoveTask(e.TaskID)
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (ia *InvoiceAggregate) CreateInvoice(userID user.UserID) error {
	i, err := NewInvoice(uuid.NewString(), userID)
	if err != nil {
		return err
	}

	ia.Invoice = i

	evt, err := NewInvoiceCreatedEvent(i.ID, i.User)
	if err != nil {
		return err
	}

	err = ia.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}

func (ia *InvoiceAggregate) AddTask(taskID TaskID) error {
	err := ia.Invoice.AddTask(taskID)
	if err != nil {
		return err
	}

	evt, err := NewTaskAddedEvent(ia.Invoice.ID, ia.Invoice.User, taskID)
	if err != nil {
		return err
	}

	err = ia.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}

func (ia *InvoiceAggregate) RemoveTask(taskID TaskID) error {
	err := ia.Invoice.RemoveTask(taskID)
	if err != nil {
		return err
	}

	evt, err := NewTaskRemovedEvent(ia.Invoice.ID, ia.Invoice.User, taskID)
	if err != nil {
		return err
	}

	err = ia.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}
