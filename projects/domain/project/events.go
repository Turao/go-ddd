package project

import (
	"errors"

	"github.com/turao/go-ddd/events"
)

type projectEventHandler struct {
	project *Project
}

func (pev *projectEventHandler) Handle(e events.DomainEvent) error {
	switch event := e.(type) {
	case ProjectCreatedEvent:
		p, err := NewProject(event.AggregateID(), event.name, true)
		if err != nil {
			return err
		}
		pev.project = p
		return nil
	case ProjectUpdatedEvent:
		pev.project.SetName(event.name)
		return nil
	case ProjectDeletedEvent:
		pev.project.Delete()
		return nil
	default:
		return nil
	}
}

// -- Events --
type ProjectCreatedEvent struct {
	events.DomainEvent
	name string
}

func NewProjectCreatedEvent(id ProjectID, name string) (events.DomainEvent, error) {
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

func NewProjectUpdatedEvent(id ProjectID, name string) (events.DomainEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.updated", id)
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

type ProjectDeletedEvent struct {
	events.DomainEvent
}

func NewProjectDeletedEvent(id ProjectID) (events.DomainEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.deleted", id)
	if err != nil {
		return nil, err
	}

	return &ProjectDeletedEvent{
		domainEvent,
	}, nil
}
