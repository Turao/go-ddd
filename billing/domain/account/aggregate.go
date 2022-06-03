package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type AccountAggregate struct {
	Account *Account `json:"account"`
	EventFactory
}

func NewAccountAggregate(ef EventFactory) *AccountAggregate {
	return &AccountAggregate{
		Account:      nil,
		EventFactory: ef,
	}
}

func (agg AccountAggregate) ID() string {
	return agg.Account.ID
}

func (agg *AccountAggregate) HandleEvent(ctx context.Context, event ddd.DomainEvent) error {
	switch e := event.(type) {
	case AccountCreatedEvent:
		a, err := NewAccount(e.AggregateID(), e.UserID, e.InvoiceID)
		if err != nil {
			return err
		}
		agg.Account = a
		return nil
	case TaskAddedEvent:
		err := agg.Account.Invoice.AddTask(e.TaskID)
		if err != nil {
			return err
		}
		return nil
	case TaskRemovedEvent:
		err := agg.Account.Invoice.RemoveTask(e.TaskID)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (agg *AccountAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case CreateAccountCommand:
		return agg.CreateAccount(c)
	case AddTaskToUserCommand:
		return agg.AddTask(c)
	case RemoveTaskFromUserCommand:
		return agg.RemoveTask(c)
	default:
		return nil, fmt.Errorf("unable to handle command %s", cmd)
	}
}

func (agg *AccountAggregate) CreateAccount(cmd CreateAccountCommand) ([]ddd.DomainEvent, error) {
	a, err := NewAccount(cmd.UserID, cmd.UserID, uuid.NewString()) // use UserID as AccountID
	if err != nil {
		return nil, err
	}
	agg.Account = a

	evt, err := agg.EventFactory.NewAccountCreatedEvent(a.ID, a.User.ID, a.Invoice.ID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (agg *AccountAggregate) assertAccountExists() error {
	if agg.Account == nil {
		return errors.New("account has not been created yet")
	}
	return nil
}

func (agg *AccountAggregate) AddTask(cmd AddTaskToUserCommand) ([]ddd.DomainEvent, error) {
	if err := agg.assertAccountExists(); err != nil {
		return nil, err
	}

	err := agg.Account.Invoice.AddTask(cmd.TaskID)
	if err != nil {
		return nil, err
	}

	evt, err := agg.EventFactory.NewTaskAddedEvent(agg.Account.ID, agg.Account.Invoice.ID, cmd.TaskID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (agg *AccountAggregate) RemoveTask(cmd RemoveTaskFromUserCommand) ([]ddd.DomainEvent, error) {
	if err := agg.assertAccountExists(); err != nil {
		return nil, err
	}

	err := agg.Account.Invoice.RemoveTask(cmd.TaskID)
	if err != nil {
		return nil, err
	}

	evt, err := agg.EventFactory.NewTaskRemovedEvent(agg.Account.ID, agg.Account.Invoice.ID, cmd.TaskID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}
