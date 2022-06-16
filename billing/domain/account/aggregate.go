package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

const AccountAggregateName = "account"

type AccountAggregate struct {
	id string

	User    *User    `json:"user"`
	Invoice *Invoice `json:"invoice"`

	EventFactory
}

var _ ddd.Aggregate = (*AccountAggregate)(nil)

var (
	ErrUnknownEvent   = errors.New("unknown event")
	ErrUnknownCommand = errors.New("unknown command")
)

type AccountAggregateOption = func(agg *AccountAggregate) error

func WithAggregateID(id string) AccountAggregateOption {
	return func(agg *AccountAggregate) error {
		if id == "" {
			return errors.New("account aggregate id is empty")
		}
		agg.id = id
		return nil
	}
}

func NewAggregate(ef EventFactory, opts ...AccountAggregateOption) (*AccountAggregate, error) {
	agg := &AccountAggregate{
		id:           uuid.NewString(),
		User:         nil,
		Invoice:      nil,
		EventFactory: ef,
	}

	for _, opt := range opts {
		if err := opt(agg); err != nil {
			return nil, err
		}
	}

	return agg, nil
}

func (agg AccountAggregate) ID() string {
	return agg.id
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
	u, err := NewUser(evt.UserID)
	if err != nil {
		return err
	}
	agg.User = u

	i, err := NewInvoice(WithInvoiceID(evt.InvoiceID))
	if err != nil {
		return err
	}
	agg.Invoice = i

	return nil
}

func (agg *AccountAggregate) handleTaskAddedEvent(evt TaskAddedEvent) error {
	err := agg.Invoice.AddTask(evt.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (agg *AccountAggregate) handleTaskRemovedEvent(evt TaskRemovedEvent) error {
	err := agg.Invoice.RemoveTask(evt.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (agg *AccountAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case CreateAccountCommand:
		return agg.handleCreateAccountCommand(c)
	case AddTaskCommand:
		return agg.handleAddTaskCommand(c)
	case RemoveTaskCommand:
		return agg.handleRemoveTaskCommand(c)
	default:
		return nil, ErrUnknownCommand
	}
}

func (agg *AccountAggregate) handleCreateAccountCommand(cmd CreateAccountCommand) ([]ddd.DomainEvent, error) {
	u, err := NewUser(cmd.UserID)
	if err != nil {
		return nil, err
	}
	agg.User = u

	i, err := NewInvoice()
	if err != nil {
		return nil, err
	}
	agg.Invoice = i

	evt, err := agg.EventFactory.NewAccountCreatedEvent(
		agg.id,
		agg.User.ID,
		agg.Invoice.ID,
	)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (agg *AccountAggregate) handleAddTaskCommand(cmd AddTaskCommand) ([]ddd.DomainEvent, error) {
	err := agg.Invoice.AddTask(cmd.TaskID)
	if err != nil {
		return nil, err
	}

	evt, err := agg.EventFactory.NewTaskAddedEvent(agg.id, agg.Invoice.ID, cmd.TaskID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (agg *AccountAggregate) handleRemoveTaskCommand(cmd RemoveTaskCommand) ([]ddd.DomainEvent, error) {
	err := agg.Invoice.RemoveTask(cmd.TaskID)
	if err != nil {
		return nil, err
	}

	evt, err := agg.EventFactory.NewTaskRemovedEvent(agg.id, agg.Invoice.ID, cmd.TaskID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (agg AccountAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID      string   `json:"id"`
		User    *User    `json:"user"`
		Invoice *Invoice `json:"invoice"`
	}{
		ID:      agg.id,
		User:    agg.User,
		Invoice: agg.Invoice,
	})
}

func (agg *AccountAggregate) UnmarshalJSON(data []byte) error {
	var payload struct {
		ID      string   `json:"id"`
		User    *User    `json:"user"`
		Invoice *Invoice `json:"invoice"`
	}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	agg.id = payload.ID
	agg.User = payload.User
	agg.Invoice = payload.Invoice
	return nil
}
