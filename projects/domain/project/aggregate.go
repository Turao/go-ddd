package project

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/domain/user"
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
		p, err := NewProject(event.AggregateID(), event.ProjectName, event.CreatedBy, event.CreatedAt, true)
		if err != nil {
			return err
		}
		pa.Project = p
		return nil
	case ProjectUpdatedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		return pa.Project.Rename(event.ProjectName)
	case ProjectDeletedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		return pa.Project.Delete()
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (pa *ProjectAggregate) CreateProject(name string, createdBy user.UserID) error {
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

func (pa *ProjectAggregate) UpdateProject(name string) error {
	err := pa.Project.Rename(name)
	if err != nil {
		return err
	}

	evt, err := NewProjectUpdatedEvent(pa.Project.ID, name)
	if err != nil {
		return err
	}

	return pa.events.Push(context.Background(), *evt)
}
