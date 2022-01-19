package task

import (
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
	"github.com/turao/go-ddd/users/domain/user"
)

type TaskID = string

type Task struct {
	ID          TaskID            `json:"id"`
	ProjectID   project.ProjectID `json:"projectId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`

	AssignedUser *user.UserID `json:"assignedUser"`
}

var (
	ErrInvalidTaskID      = errors.New("invalid task id")
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidUserID      = errors.New("invalid user id")
)

func NewTask(id TaskID, projectId project.ProjectID, title string, description string) (*Task, error) {
	if id == "" {
		return nil, ErrInvalidTaskID
	}

	if projectId == "" {
		return nil, ErrInvalidProjectID
	}

	if title == "" {
		return nil, ErrInvalidTitle
	}

	if description == "" {
		return nil, ErrInvalidDescription
	}

	return &Task{
		ID:           id,
		ProjectID:    projectId,
		Title:        title,
		Description:  description,
		AssignedUser: nil,
	}, nil
}

func (t *Task) AssignTo(assignedUserID user.UserID) error {
	t.AssignedUser = &assignedUserID
	return nil
}
