package task

import (
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
)

type TaskID = string

type Task struct {
	ID          TaskID            `json:"id"`
	ProjectID   project.ProjectID `json:"projectId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
}

var (
	ErrInvalidTaskID      = errors.New("invalid task id")
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidDescription = errors.New("invalid description")
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
		ID: id,
	}, nil
}
