package project

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type ProjectAggregate struct {
	Project *Project `json:"project"`

	events events.EventStore
}

func NewProjectAggregate(p *Project, es events.EventStore) (*ProjectAggregate, error) {
	return &ProjectAggregate{
		Project: p,
		events:  es,
	}, nil
}

func (pa *ProjectAggregate) HandleEvent(e events.DomainEvent) error {
	switch event := e.(type) {
	case ProjectCreatedEvent:
		p, err := NewProject(event.AggregateID(), event.ProjectName, true)
		if err != nil {
			return err
		}
		pa.Project = p
		return nil
	case ProjectUpdatedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		pa.Project.Rename(event.ProjectName)
		return nil
	case ProjectDeletedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		pa.Project.Delete()
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (pa ProjectAggregate) CreateProject(name string) error {
	p, err := NewProject(uuid.NewString(), name, true)
	if err != nil {
		return err
	}

	evt, err := NewProjectCreatedEvent(p.ID, p.Name)
	if err != nil {
		return err
	}

	err = pa.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}

func (pa *ProjectAggregate) DeleteProject() error {
	err := pa.Project.Delete()
	if err != nil {
		return err
	}

	evt, err := NewProjectDeletedEvent(pa.Project.ID)
	if err != nil {
		return err
	}

	return pa.events.Push(context.Background(), *evt)
}
