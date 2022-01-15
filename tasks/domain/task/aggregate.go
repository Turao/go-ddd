package task

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type TaskAggregate struct {
	Task *Task
}

func (ta TaskAggregate) HandleEvent(e events.DomainEvent) error {
	switch event := e.(type) {
	case TaskCreatedEvent:
		t, err := NewTask(event.AggregateID(), event.ProjectID, event.Title, event.Description)
		if err != nil {
			return err
		}
		ta.Task = t
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func CreateTask(projectID project.ProjectID, title string, description string) (*Task, error) {
	return NewTask(uuid.NewString(), projectID, title, description)
}
