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
		if pa.Project == nil {
			return errors.New("project does not exist")
		}
		pa.Project.Rename(event.projectName)
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