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

type TaskUnassignedEvent struct {
	events.DomainEvent `json:"domainEvent"`
}

func NewTaskUnassignedEvent(id TaskID) (*TaskAssignedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.unassigned", id)
	if err != nil {
		return nil, err
	}

	return &TaskAssignedEvent{
		DomainEvent: domainEvent,
	}, nil
}

type TitleUpdatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	Title              string `json:"title"`
}

func NewTitleUpdatedEvent(id TaskID, title string) (*TitleUpdatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.title.updated", id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, ErrInvalidTitle
	}

	return &TitleUpdatedEvent{
		DomainEvent: domainEvent,
		Title:       title,
	}, nil
}

type DescriptionUpdatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	Description        string `json:"description"`
}

func NewDescriptionUpdatedEvent(id TaskID, description string) (*DescriptionUpdatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.description.updated", id)
	if err != nil {
		return nil, err
	}

	if description == "" {
		return nil, ErrInvalidDescription
	}

	return &DescriptionUpdatedEvent{
		DomainEvent: domainEvent,
		Description: description,
	}, nil
}
