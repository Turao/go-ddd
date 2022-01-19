package task

import (
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
	"github.com/turao/go-ddd/users/domain/user"
)

type TaskCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	ProjectID          project.ProjectID `json:"projectId"`
	Title              string            `json:"title"`
	Description        string            `json:"description"`
}

// var (
// ErrInvalidProjectID = errors.New("invalid project id")
// ErrInvalidTitle       = errors.New("invalid title")
// ErrInvalidDescription = errors.New("invalid description")
// )

func NewTaskCreatedEvent(id TaskID, projectID project.ProjectID, title string, description string) (*TaskCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.created", id)
	if err != nil {
		return nil, err
	}

	if projectID == "" {
		return nil, ErrInvalidProjectID
	}

	if title == "" {
		return nil, ErrInvalidTitle
	}

	if description == "" {
		return nil, ErrInvalidDescription
	}

	return &TaskCreatedEvent{
		DomainEvent: domainEvent,
		ProjectID:   projectID,
		Title:       title,
		Description: description,
	}, nil
}

type TaskAssignedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	AssignedTo         user.UserID `json:"assignedTo"`
}

func NewTaskAssignedEvent(id TaskID, assignedUserID user.UserID) (*TaskAssignedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.assigned", id)
	if err != nil {
		return nil, err
	}

	if assignedUserID == "" {
		return nil, ErrInvalidUserID
	}

	return &TaskAssignedEvent{
		DomainEvent: domainEvent,
		AssignedTo:  assignedUserID,
	}, nil
}
