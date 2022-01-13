package task

import (
	"fmt"

	"github.com/turao/go-ddd/events"
)

type TaskAggregate struct {
	Task *Task
}

func (ta TaskAggregate) HandleEvent(e events.DomainEvent) error {
	switch event := e.(type) {
	case TaskCreatedEvent:
		t, err := NewTask(event.AggregateID(), event.ProjectID)
		if err != nil {
			return err
		}
		ta.Task = t
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}
