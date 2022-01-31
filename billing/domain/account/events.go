package account

import "github.com/turao/go-ddd/events"

type AccountCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	UserID             UserID    `json:"userId"`
	InvoiceID          InvoiceID `json:"invoiceID"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidTaskID   = errors.New("invalid task id")
// )

func NewAccountCreatedEvent(accountID AccountID, userID UserID, invoiceID InvoiceID) (*AccountCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("account.created", accountID)
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
	events.DomainEvent `json:"domainEvent"`
	InvoiceID          InvoiceID `json:"invoiceId"`
	TaskID             TaskID    `json:"taskId"`
}

func NewTaskAddedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskAddedEvent, error) {
	domainEvent, err := events.NewDomainEvent("account.invoice.task.added", accountID)
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
	events.DomainEvent `json:"domainEvent"`
	InvoiceID          InvoiceID `json:"invoiceId"`
	TaskID             TaskID    `json:"taskId"`
}

func NewTaskRemovedEvent(accountID AccountID, invoiceID InvoiceID, taskID TaskID) (*TaskRemovedEvent, error) {
	domainEvent, err := events.NewDomainEvent("account.invoice.task.removed", accountID)
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
