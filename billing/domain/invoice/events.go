package invoice

import "github.com/turao/go-ddd/events"

type InvoiceCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	UserID             UserID `json:"userId"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidTaskID   = errors.New("invalid task id")
// )

func NewInvoiceCreatedEvent(invoiceID string, userID string) (*InvoiceCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("invoice.created", invoiceID)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &InvoiceCreatedEvent{
		DomainEvent: domainEvent,
		UserID:      userID,
	}, nil
}

type TaskAddedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	UserID             UserID `json:"userId"`
	TaskID             TaskID `json:"taskId"`
}

func NewTaskAddedEvent(invoiceID string, userID string, taskID string) (*TaskAddedEvent, error) {
	domainEvent, err := events.NewDomainEvent("invoice.task.added", invoiceID)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskAddedEvent{
		DomainEvent: domainEvent,
		UserID:      userID,
		TaskID:      taskID,
	}, nil
}

type TaskRemovedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	UserID             UserID `json:"userId"`
	TaskID             TaskID `json:"taskId"`
}

func NewTaskRemovedEvent(invoiceID string, userID string, taskID string) (*TaskRemovedEvent, error) {
	domainEvent, err := events.NewDomainEvent("invoice.task.removed", invoiceID)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskRemovedEvent{
		DomainEvent: domainEvent,
		UserID:      userID,
		TaskID:      taskID,
	}, nil
}
