package task

import (
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type EventFactory interface {
	NewTaskCreatedEvent(id TaskID, projectID ProjectID, title string, description string) (*TaskCreatedEvent, error)
	NewTaskAssignedEvent(id TaskID, assignedUserID UserID) (*TaskAssignedEvent, error)
	NewTaskUnassignedEvent(id TaskID) (*TaskAssignedEvent, error)
	NewTitleUpdatedEvent(id TaskID, title string) (*TitleUpdatedEvent, error)
	NewDescriptionUpdatedEvent(id TaskID, description string) (*DescriptionUpdatedEvent, error)
	NewStatusUpdatedEvent(id TaskID, status string) (*StatusUpdatedEvent, error)
}

type TaskEventFactory struct{}

type TaskCreatedEvent struct {
	ddd.DomainEvent `json:"domainEvent"`
	ProjectID       ProjectID `json:"projectId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
}

func (f TaskEventFactory) NewTaskCreatedEvent(id TaskID, projectID ProjectID, title string, description string) (*TaskCreatedEvent, error) {
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

func (f TaskEventFactory) NewTaskAssignedEvent(id TaskID, assignedUserID UserID) (*TaskAssignedEvent, error) {
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

func (f TaskEventFactory) NewTaskUnassignedEvent(id TaskID) (*TaskAssignedEvent, error) {
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

func (f TaskEventFactory) NewTitleUpdatedEvent(id TaskID, title string) (*TitleUpdatedEvent, error) {
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

func (f TaskEventFactory) NewDescriptionUpdatedEvent(id TaskID, description string) (*DescriptionUpdatedEvent, error) {
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

func (f TaskEventFactory) NewStatusUpdatedEvent(id TaskID, status string) (*StatusUpdatedEvent, error) {
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
