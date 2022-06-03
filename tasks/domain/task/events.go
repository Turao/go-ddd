package task

import (
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type TaskCreatedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	ProjectID       ProjectID `json:"projectId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
}

// var (
// ErrInvalidProjectID = errors.New("invalid project id")
// ErrInvalidTitle       = errors.New("invalid title")
// ErrInvalidDescription = errors.New("invalid description")
// )

func NewTaskCreatedEvent(id TaskID, projectID ProjectID, title string, description string) (*TaskCreatedEvent, error) {
	event, err := events.NewEvent("task.created")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
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
	ddd.DomainEvent `json:"domainEvent"`
	AssignedTo      UserID `json:"assignedTo"`
}

func NewTaskAssignedEvent(id TaskID, assignedUserID UserID) (*TaskAssignedEvent, error) {
	event, err := events.NewEvent("task.assigned")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
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
	ddd.DomainEvent `json:"domainEvent"`
}

func NewTaskUnassignedEvent(id TaskID) (*TaskAssignedEvent, error) {
	event, err := events.NewEvent("task.unassigned")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
	if err != nil {
		return nil, err
	}

	return &TaskAssignedEvent{
		DomainEvent: domainEvent,
	}, nil
}

type TitleUpdatedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	Title           string `json:"title"`
}

func NewTitleUpdatedEvent(id TaskID, title string) (*TitleUpdatedEvent, error) {
	event, err := events.NewEvent("task.title.updated")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
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
	ddd.DomainEvent `json:"domainEvent"`
	Description     string `json:"description"`
}

func NewDescriptionUpdatedEvent(id TaskID, description string) (*DescriptionUpdatedEvent, error) {
	event, err := events.NewEvent("task.description.updated")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
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

type StatusUpdatedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	Status          string `json:"status"`
}

func NewStatusUpdatedEvent(id TaskID, status string) (*StatusUpdatedEvent, error) {
	event, err := events.NewEvent("task.status.updated")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
	if err != nil {
		return nil, err
	}

	if status == "" {
		return nil, ErrInvalidStatus
	}

	return &StatusUpdatedEvent{
		DomainEvent: domainEvent,
		Status:      status,
	}, nil
}
