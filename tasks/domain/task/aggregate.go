package task

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type TaskAggregate struct {
	Task    *Task
	version int
	events  events.EventStore
}

func NewTaskAggregate(task *Task, es events.EventStore) (*TaskAggregate, error) {
	return &TaskAggregate{
		Task:    task,
		version: 0,
		events:  es,
	}, nil
}

func (ta *TaskAggregate) HandleEvent(e ddd.DomainEvent) error {
	switch event := e.(type) {
	case TaskCreatedEvent:
		t, err := NewTask(event.AggregateID(), event.ProjectID, event.Title, event.Description)
		if err != nil {
			return err
		}
		ta.Task = t
		ta.version += 1
		return nil
	case TaskAssignedEvent:
		err := ta.Task.AssignTo(event.AssignedTo)
		if err != nil {
			return err
		}
		ta.version += 1
		return nil
	case TaskUnassignedEvent:
		err := ta.Task.Unassign()
		if err != nil {
			return err
		}
		ta.version += 1
		return nil
	case TitleUpdatedEvent:
		err := ta.Task.UpdateTitle(event.Title)
		if err != nil {
			return err
		}
		ta.version += 1
		return nil
	case DescriptionUpdatedEvent:
		err := ta.Task.UpdateDescription(event.Description)
		if err != nil {
			return err
		}
		ta.version += 1
		return nil
	case StatusUpdatedEvent:
		err := ta.Task.UpdateStatus(event.Status)
		if err != nil {
			return err
		}
		ta.version += 1
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (ta *TaskAggregate) CreateTask(projectID ProjectID, title string, description string) (*Task, error) {
	t, err := NewTask(uuid.NewString(), projectID, title, description)
	if err != nil {
		return nil, err
	}

	ta.Task = t

	evt, err := NewTaskCreatedEvent(t.ID, t.ProjectID, t.Title, t.Description)
	if err != nil {
		return nil, err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return nil, err
	}
	ta.version += 1

	return t, nil
}

func (ta TaskAggregate) AssignTo(assignedUserID UserID) error {
	err := ta.Task.AssignTo(assignedUserID)
	if err != nil {
		return err
	}

	evt, err := NewTaskAssignedEvent(ta.Task.ID, assignedUserID)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return err
	}
	ta.version += 1

	return nil
}

func (ta TaskAggregate) Unassign() error {
	err := ta.Task.Unassign()
	if err != nil {
		return err
	}

	evt, err := NewTaskUnassignedEvent(ta.Task.ID)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return err
	}
	ta.version += 1

	return nil
}

func (ta TaskAggregate) UpdateTitle(title string) error {
	err := ta.Task.UpdateTitle(title)
	if err != nil {
		return err
	}

	evt, err := NewTitleUpdatedEvent(ta.Task.ID, title)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return err
	}
	ta.version += 1

	return nil
}

func (ta TaskAggregate) UpdateDescription(description string) error {
	err := ta.Task.UpdateDescription(description)
	if err != nil {
		return err
	}

	evt, err := NewDescriptionUpdatedEvent(ta.Task.ID, description)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return err
	}
	ta.version += 1

	return nil
}

func (ta TaskAggregate) UpdateStatus(status string) error {
	err := ta.Task.UpdateStatus(status)
	if err != nil {
		return err
	}

	evt, err := NewStatusUpdatedEvent(ta.Task.ID, status)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt, ta.version+1)
	if err != nil {
		return err
	}
	ta.version += 1

	return nil
}
