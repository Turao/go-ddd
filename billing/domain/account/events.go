package account

import (
	"github.com/turao/go-ddd/ddd"
	v1 "github.com/turao/go-ddd/events/v1"
)

type EventFactory interface {
	NewAccountCreatedEvent(accountID AccountID, userID UserID, invoiceID InvoiceID) (*AccountCreatedEvent, error)
	NewTaskAddedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskAddedEvent, error)
	NewTaskRemovedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskRemovedEvent, error)
}

type AccountEventsFactory struct{}

type AccountCreatedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	UserID          UserID    `json:"userId"`
	InvoiceID       InvoiceID `json:"invoiceID"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidTaskID   = errors.New("invalid task id")
// )

func (f AccountEventsFactory) NewAccountCreatedEvent(accountID AccountID, userID UserID, invoiceID InvoiceID) (*AccountCreatedEvent, error) {
	event, err := v1.NewEvent("account.created")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, accountID, AccountAggregateName)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	if invoiceID == "" {
		return nil, ErrInvalidInvoiceID
	}

	return &AccountCreatedEvent{
		DomainEvent: domainEvent,
		UserID:      userID,
		InvoiceID:   invoiceID,
	}, nil
}

type TaskAddedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	InvoiceID       InvoiceID `json:"invoiceId"`
	TaskID          TaskID    `json:"taskId"`
}

func (f AccountEventsFactory) NewTaskAddedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskAddedEvent, error) {
	event, err := v1.NewEvent("account.invoice.task.added")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, accountID, AccountAggregateName)
	if err != nil {
		return nil, err
	}

	if invoiceID == "" {
		return nil, ErrInvalidInvoiceID
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskAddedEvent{
		DomainEvent: domainEvent,
		InvoiceID:   invoiceID,
		TaskID:      taskID,
	}, nil
}

type TaskRemovedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	InvoiceID       InvoiceID `json:"invoiceId"`
	TaskID          TaskID    `json:"taskId"`
}

func (f AccountEventsFactory) NewTaskRemovedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskRemovedEvent, error) {
	event, err := v1.NewEvent("account.invoice.task.removed")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, accountID, AccountAggregateName)
	if err != nil {
		return nil, err
	}

	if invoiceID == "" {
		return nil, ErrInvalidInvoiceID
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskRemovedEvent{
		DomainEvent: domainEvent,
		InvoiceID:   invoiceID,
		TaskID:      taskID,
	}, nil
}
