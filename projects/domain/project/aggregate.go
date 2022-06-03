package project

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type ProjectAggregate struct {
	Project *Project `json:"project"`
	version int
	events  events.EventStore
}

func NewProjectAggregate(p *Project, es events.EventStore) (*ProjectAggregate, error) {
	return &ProjectAggregate{
		Project: p,
		version: 0,
		events:  es,
	}, nil
}

func (pa *ProjectAggregate) HandleEvent(e ddd.DomainEvent) error {
	switch event := e.(type) {
	case ProjectCreatedEvent:
		p, err := NewProject(event.AggregateID(), event.ProjectName, event.CreatedBy, event.CreatedAt, true)
		if err != nil {
			return err
		}
		pa.Project = p
		pa.version += 1
		return nil
	case ProjectUpdatedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		err := pa.Project.Rename(event.ProjectName)
		if err != nil {
			return err
		}
		pa.version += 1
		return nil
	case ProjectDeletedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		err := pa.Project.Delete()
		if err != nil {
			return err
		}
		pa.version += 1
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (pa *ProjectAggregate) CreateProject(name string, createdBy UserID) error {
	now := time.Now()
	p, err := NewProject(uuid.NewString(), name, createdBy, now, true)
	if err != nil {
		return err
	}

	pa.Project = p

	evt, err := NewProjectCreatedEvent(p.ID, p.Name, createdBy, now)
	if err != nil {
		return err
	}

	err = pa.events.Push(context.Background(), *evt, pa.version+1)
	if err != nil {
		return err
	}
	pa.version += 1

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

	err = pa.events.Push(context.Background(), *evt, pa.version+1)
	if err != nil {
		return err
	}
	pa.version += 1

	return nil
}

func (pa *ProjectAggregate) UpdateProject(name string) error {
	err := pa.Project.Rename(name)
	if err != nil {
		return err
	}

	evt, err := NewProjectUpdatedEvent(pa.Project.ID, name)
	if err != nil {
		return err
	}

	err = pa.events.Push(context.Background(), *evt, pa.version+1)
	if err != nil {
		return err
	}
	pa.version += 1

	return nil
}
