package project

import (
	"errors"
	"fmt"

	"github.com/turao/go-ddd/events"
)

type ProjectAggregate struct {
	Project *Project `json:"project"`
}

func (pa *ProjectAggregate) HandleEvents(es []events.DomainEvent) error {
	for _, e := range es {
		if err := pa.HandleEvent(e); err != nil {
			return err
		}
	}
	return nil
}

func (pa *ProjectAggregate) HandleEvent(e events.DomainEvent) error {
	switch event := e.(type) {
	case ProjectCreatedEvent:
		p, err := NewProject(event.AggregateID(), event.projectName, true)
		if err != nil {
			return err
		}
		pa.Project = p
		return nil
	case ProjectUpdatedEvent:
		pa.Project.SetName(event.projectName)
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
	events.DomainEvent `json:"domainEvent"`
	projectName        string `json:"projectName"`
}

func NewProjectCreatedEvent(id ProjectID, projectName string) (*ProjectCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.created", id)
	if err != nil {
		return nil, err
	}

	if projectName == "" {
		return nil, errors.New("project name must not be empty")
	}

	return &ProjectCreatedEvent{
		domainEvent,
		projectName,
	}, nil
}

type ProjectUpdatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	projectName        string `json:"projectName"`
}

func NewProjectUpdatedEvent(id ProjectID, projectName string) (*ProjectUpdatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.updated", id)
	if err != nil {
		return nil, err
	}

	if projectName == "" {
		return nil, errors.New("project name must not be empty")
	}

	return &ProjectUpdatedEvent{
		domainEvent,
		projectName,
	}, nil
}

type ProjectDeletedEvent struct {
	events.DomainEvent `json:"domainEvent"`
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
