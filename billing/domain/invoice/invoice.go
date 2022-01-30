package invoice

import (
	"errors"

	"github.com/turao/go-ddd/tasks/domain/task"
	"github.com/turao/go-ddd/users/domain/user"
)

type InvoiceID = string
type UserID = user.UserID
type TaskID = task.TaskID

type Invoice struct {
	ID             InvoiceID       `json:"id"`
	User           UserID          `json:"userId"`
	OpenTasks      map[TaskID]bool `json:"openTasks"`
	CompletedTasks map[TaskID]bool `json:"completedTasks"`
}

var (
	ErrInvalidInvoiceID = errors.New("invalid invoice id")
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidTaskID    = errors.New("invalid task id")
)

func NewInvoice(invoiceID InvoiceID, userID UserID) (*Invoice, error) {
	if invoiceID == "" {
		return nil, ErrInvalidInvoiceID
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &Invoice{
		ID:             invoiceID,
		User:           userID,
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
