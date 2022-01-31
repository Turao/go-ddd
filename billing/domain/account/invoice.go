package account

import (
	"errors"

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

func NewInvoice(invoiceID InvoiceID) (*Invoice, error) {
	if invoiceID == "" {
		return nil, ErrInvalidInvoiceID
	}

	return &Invoice{
		ID:             invoiceID,
		OpenTasks:      make(map[TaskID]bool),
		CompletedTasks: make(map[TaskID]bool),
	}, nil
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
