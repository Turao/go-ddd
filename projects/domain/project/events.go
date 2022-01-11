package project

import (
	"errors"
	"fmt"

	"github.com/turao/go-ddd/events"
)

type ProjectAggregate struct {
	Project *Project
}

func (pa *ProjectAggregate) Handle(e events.DomainEvent) error {
	switch event := e.(type) {
	case ProjectCreatedEvent:
		p, err := NewProject(event.AggregateID(), event.name, true)
		if err != nil {
			return err
		}
		pa.Project = p
		return nil
	case ProjectUpdatedEvent:
		pa.Project.SetName(event.name)
		return nil
	case ProjectDeletedEvent:
		pa.Project.Delete()
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

// -- Events --
type ProjectCreatedEvent struct {
	events.DomainEvent
	name string
}

func NewProjectCreatedEvent(id ProjectID, name string) (*ProjectCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.created", id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &ProjectCreatedEvent{
		domainEvent,
		name,
	}, nil
}

type ProjectUpdatedEvent struct {
	events.DomainEvent
	name string
}

func NewProjectUpdatedEvent(id ProjectID, name string) (*ProjectUpdatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.updated", id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	return &ProjectUpdatedEvent{
		domainEvent,
		name,
	}, nil
}

type ProjectDeletedEvent struct {
	events.DomainEvent
}

func NewProjectDeletedEvent(id ProjectID) (*ProjectDeletedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.deleted", id)
	if err != nil {
		return nil, err
	}

	return &ProjectDeletedEvent{
		domainEvent,
	}, nil
}
