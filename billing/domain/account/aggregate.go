package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/domain/user"
)

type AccountAggregate struct {
	Account *Account `json:"account"`
	events  events.EventStore
}

func NewAccountAggregate(a *Account, es events.EventStore) (*AccountAggregate, error) {
	return &AccountAggregate{
		Account: a,
		events:  es,
	}, nil
}

func (aa *AccountAggregate) HandleEvent(event events.DomainEvent) error {
	switch e := event.(type) {
	case AccountCreatedEvent:
		a, err := NewAccount(e.AggregateID(), e.UserID, e.InvoiceID)
		if err != nil {
			return err
		}
		aa.Account = a
		return nil
	case TaskAddedEvent:
		return aa.Account.Invoice.AddTask(e.TaskID)
	case TaskRemovedEvent:
		return aa.Account.Invoice.RemoveTask(e.TaskID)
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (aa *AccountAggregate) CreateAccount(userID user.UserID) error {
	a, err := NewAccount(userID, userID, uuid.NewString()) // use UserID as AccountID
	if err != nil {
		return err
	}
	aa.Account = a

	evt, err := NewAccountCreatedEvent(a.ID, a.User.ID, a.Invoice.ID)
	if err != nil {
		return err
	}

	err = aa.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}

func (aa *AccountAggregate) assertAccountExists() error {
	if aa.Account == nil {
		return errors.New("account has not been created yet")
	}
	return nil
}

func (aa *AccountAggregate) AddTask(taskID TaskID) error {
	if err := aa.assertAccountExists(); err != nil {
		return err
	}

	err := aa.Account.Invoice.AddTask(taskID)
	if err != nil {
		return err
	}

	evt, err := NewTaskAddedEvent(aa.Account.ID, aa.Account.Invoice.ID, taskID)
	if err != nil {
		return err
	}

	err = aa.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}

func (aa *AccountAggregate) RemoveTask(taskID TaskID) error {
	if err := aa.assertAccountExists(); err != nil {
		return err
	}

	err := aa.Account.Invoice.RemoveTask(taskID)
	if err != nil {
		return err
	}

	evt, err := NewTaskRemovedEvent(aa.Account.ID, aa.Account.Invoice.ID, taskID)
	if err != nil {
		return err
	}

	err = aa.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}
