package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type AccountAggregate struct {
	Account *Account `json:"account"`
	EventFactory
}

var _ ddd.Aggregate = (*AccountAggregate)(nil)

var (
	ErrUnknownEvent   = errors.New("unknown event")
	ErrUnknownCommand = errors.New("unknown command")
)

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
		return agg.handleAccountCreatedEvent(e)
	case TaskAddedEvent:
		return agg.handleTaskAddedEvent(e)
	case TaskRemovedEvent:
		return agg.handleTaskRemovedEvent(e)
	default:
		return ErrUnknownEvent
	}
}

func (agg *AccountAggregate) handleAccountCreatedEvent(evt AccountCreatedEvent) error {
	a, err := NewAccount(evt.AggregateID(), evt.UserID, evt.InvoiceID)
	if err != nil {
		return err
	}
	agg.Account = a
	return nil
}

func (agg *AccountAggregate) handleTaskAddedEvent(evt TaskAddedEvent) error {
	err := agg.Account.Invoice.AddTask(evt.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (agg *AccountAggregate) handleTaskRemovedEvent(evt TaskRemovedEvent) error {
	err := agg.Account.Invoice.RemoveTask(evt.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (agg *AccountAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case CreateAccountCommand:
		return agg.handleCreateAccountCommand(c)
	case AddTaskToUserCommand:
		return agg.handleAddTaskCommand(c)
	case RemoveTaskFromUserCommand:
		return agg.handleRemoveTaskCommand(c)
	default:
		return nil, ErrUnknownCommand
	}
}

func (agg *AccountAggregate) handleCreateAccountCommand(cmd CreateAccountCommand) ([]ddd.DomainEvent, error) {
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

func (agg *AccountAggregate) handleAddTaskCommand(cmd AddTaskToUserCommand) ([]ddd.DomainEvent, error) {
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

func (agg *AccountAggregate) handleRemoveTaskCommand(cmd RemoveTaskFromUserCommand) ([]ddd.DomainEvent, error) {
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

func (agg AccountAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Account Account `json:"account"`
	}{
		Account: *agg.Account,
	})
}

func (agg *AccountAggregate) UnmarshalJSON(data []byte) error {
	var payload struct {
		Account Account `json:"account"`
	}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}
	agg.Account = &payload.Account
	return nil
}
