package project

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type ProjectAggregate struct {
	Project *Project `json:"project"`

	EventFactory
}

var (
	ErrUnknownEvent   = errors.New("unknown event")
	ErrUnknownCommand = errors.New("unknown command")
)

func NewProjectAggregate(ef EventFactory) *ProjectAggregate {
	return &ProjectAggregate{
		Project:      nil,
		EventFactory: ef,
	}
}

func (pa ProjectAggregate) ID() string {
	return pa.Project.ID
}

func (pa *ProjectAggregate) HandleEvent(ctx context.Context, e ddd.DomainEvent) error {
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
		err := pa.Project.Rename(event.ProjectName)
		if err != nil {
			return err
		}
		return nil
	case ProjectDeletedEvent:
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		err := pa.Project.Delete()
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrUnknownEvent
	}
}

func (pa *ProjectAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case CreateProjectCommand:
		return pa.CreateProject(c.Name, c.CreatedBy)
	case UpdateProjectCommand:
		return pa.UpdateProject(c.Name)
	case DeleteProjectCommand:
		return pa.DeleteProject()
	default:
		return nil, ErrUnknownCommand
	}
}

func (pa *ProjectAggregate) CreateProject(name string, createdBy UserID) ([]ddd.DomainEvent, error) {
	now := time.Now()
	p, err := NewProject(uuid.NewString(), name, createdBy, now, true)
	if err != nil {
		return nil, err
	}

	pa.Project = p

	evt, err := pa.EventFactory.NewProjectCreatedEvent(p.ID, p.Name, createdBy, now)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (pa *ProjectAggregate) UpdateProject(name string) ([]ddd.DomainEvent, error) {
	err := pa.Project.Rename(name)
	if err != nil {
		return nil, err
	}

	evt, err := pa.EventFactory.NewProjectUpdatedEvent(pa.Project.ID, name)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (pa *ProjectAggregate) DeleteProject() ([]ddd.DomainEvent, error) {
	err := pa.Project.Delete()
	if err != nil {
		return nil, err
	}

	evt, err := pa.EventFactory.NewProjectDeletedEvent(pa.Project.ID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}
