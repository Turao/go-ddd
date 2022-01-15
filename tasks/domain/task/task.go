package task

import (
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
)

type TaskID = string

type Task struct {
	ID        TaskID            `json:"id"`
	ProjectID project.ProjectID `json:"projectId"`
}

var (
	ErrInvalidTaskID      = errors.New("invalid task id")
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidDescription = errors.New("invalid description")
)

func NewTask(id TaskID, projectId project.ProjectID) (*Task, error) {
	if id == "" {
		return nil, ErrInvalidTaskID
	}

	if projectId == "" {
		return nil, ErrInvalidProjectID
	}

	return &Task{
		ID: id,
	}, nil
}
