package task

import (
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
	"github.com/turao/go-ddd/users/domain/user"
)

type TaskID = string
type ProjectID = project.ProjectID
type UserID = user.UserID

type Status = string

const (
	New        Status = "new"
	InProgress Status = "in_progress"
	Blocked    Status = "blocked"
	Completed  Status = "completed"
)

type Task struct {
	ID          TaskID    `json:"id"`
	ProjectID   ProjectID `json:"projectId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`

	AssignedUser *UserID `json:"assignedUser"`
}

var (
	ErrInvalidTaskID      = errors.New("invalid task id")
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidStatus      = errors.New("invalid status")
)

func NewTask(id TaskID, projectId ProjectID, title string, description string) (*Task, error) {
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
		Status:       New,
		AssignedUser: nil,
	}, nil
}

func (t *Task) AssignToUser(assignedUserID UserID) error {
	t.AssignedUser = &assignedUserID
	return nil
}

func (t *Task) Unassign() error {
	t.AssignedUser = nil
	return nil
}

func (t *Task) UpdateStatus(status Status) error {
	switch status {
	case New:
		t.Status = New
		return nil
	case InProgress:
		t.Status = InProgress
		return nil
	case Blocked:
		t.Status = Blocked
		return nil
	case Completed:
		t.Status = Completed
		return nil
	default:
		return ErrInvalidStatus
	}
}

func (t *Task) UpdateTitle(title string) error {
	if title == "" {
		return ErrInvalidTitle
	}
	t.Title = title
	return nil
}

func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return ErrInvalidDescription
	}
	t.Description = description
	return nil
}
