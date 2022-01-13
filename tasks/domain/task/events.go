package task

import (
	"errors"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type TaskCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	ProjectID          project.ProjectID `json:"projectId"`
}

func NewTaskCreatedEvent(id TaskID, projectID project.ProjectID) (*TaskCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.created", id)
	if err != nil {
		return nil, err
	}

	if projectID == "" {
		return nil, errors.New("invalid project id")
	}

	return &TaskCreatedEvent{
		DomainEvent: domainEvent,
		ProjectID:   projectID,
	}, nil
}
