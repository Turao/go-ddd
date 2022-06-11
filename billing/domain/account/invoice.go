package account

import (
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type InvoiceID = string
type TaskID = task.TaskID

type Invoice struct {
	ID             InvoiceID       `json:"id"`
	OpenTasks      map[TaskID]bool `json:"openTasks"`
	CompletedTasks map[TaskID]bool `json:"completedTasks"`
}

var (
	ErrInvalidInvoiceID = errors.New("invalid invoice id")
	ErrInvalidTaskID    = errors.New("invalid task id")
)

type InvoiceOption = func(invoice *Invoice) error

func WithInvoiceID(id InvoiceID) InvoiceOption {
	return func(invoice *Invoice) error {
		if id == "" {
			return ErrInvalidInvoiceID
		}
		invoice.ID = id
		return nil
	}
}

func NewInvoice(opts ...InvoiceOption) (*Invoice, error) {
	invoice := &Invoice{
		ID:             uuid.NewString(),
		OpenTasks:      make(map[TaskID]bool),
		CompletedTasks: make(map[TaskID]bool),
	}

	for _, opt := range opts {
		if err := opt(invoice); err != nil {
			return nil, err
		}
	}

	return invoice, nil
}

func (u *Invoice) AddTask(taskID TaskID) error {
	if taskID == "" {
		return ErrInvalidTaskID
	}

	u.OpenTasks[taskID] = true
	return nil
}

func (u *Invoice) RemoveTask(taskID TaskID) error {
	if taskID == "" {
		return ErrInvalidTaskID
	}

	delete(u.OpenTasks, taskID)
	return nil
}
