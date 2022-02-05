package account

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/shared/ddd"
	"github.com/turao/go-ddd/users/domain/user"
)

type AccountAggregate struct {
	ddd.AggregateRoot

	Account *Account `json:"account"`
}

func NewAccountAggregate(a *Account, es events.EventStore) (*AccountAggregate, error) {
	return &AccountAggregate{
		Account: a,
		AggregateRoot: ddd.AggregateRoot{
			Version: 0,
			Events:  es,
		},
	}, nil
}

func (aa *AccountAggregate) HandleEvent(event events.DomainEvent) error {
	switch e := event.(type) {
	case AccountCreatedEvent:
		return aa.AggregateRoot.HandleEvent(event, aa.OnAccountCreatedEvent)
	case TaskAddedEvent:
		return aa.AggregateRoot.HandleEvent(event, aa.OnTaskAddedEvent)
	case TaskRemovedEvent:
		return aa.AggregateRoot.HandleEvent(event, aa.OnTaskRemovedEvent)
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

	err = aa.AggregateRoot.AddEvent(*evt)
	if err != nil {
		return err
	}

	return nil
}

func (aa *AccountAggregate) OnAccountCreatedEvent(event events.DomainEvent) error {
	e := event.(AccountCreatedEvent)
	a, err := NewAccount(e.AggregateID(), e.UserID, e.InvoiceID)
	if err != nil {
		return err
	}
	aa.Account = a
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

	err = aa.AggregateRoot.AddEvent(*evt)
	if err != nil {
		return err
	}

	return nil
}

func (aa *AccountAggregate) OnTaskAddedEvent(event events.DomainEvent) error {
	e := event.(TaskAddedEvent)
	return aa.Account.Invoice.AddTask(e.TaskID)
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

	err = aa.AggregateRoot.AddEvent(*evt)
	if err != nil {
		return err
	}

	return nil
}

func (aa *AccountAggregate) OnTaskRemovedEvent(event events.DomainEvent) error {
	e := event.(TaskRemovedEvent)
	return aa.Account.Invoice.RemoveTask(e.TaskID)
}
