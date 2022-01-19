package task

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
	"github.com/turao/go-ddd/users/domain/user"
)

type TaskAggregate struct {
	Task *Task

	events events.EventStore
}

func NewTaskAggregate(task *Task, es events.EventStore) (*TaskAggregate, error) {
	return &TaskAggregate{
		Task:   task,
		events: es,
	}, nil
}

func (ta *TaskAggregate) HandleEvent(e events.DomainEvent) error {
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

func (ta *TaskAggregate) CreateTask(projectID project.ProjectID, title string, description string) (*Task, error) {
	t, err := NewTask(uuid.NewString(), projectID, title, description)
	if err != nil {
		return nil, err
	}

	ta.Task = t

	evt, err := NewTaskCreatedEvent(t.ID, t.ProjectID, t.Title, t.Description)
	if err != nil {
		return nil, err
	}

	err = ta.events.Push(context.Background(), *evt)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (ta TaskAggregate) AssignTo(assignedUserID user.UserID) error {
	err := ta.Task.AssignTo(assignedUserID)
	if err != nil {
		return err
	}

	evt, err := NewTaskAssignedEvent(ta.Task.ID, assignedUserID)
	if err != nil {
		return err
	}

	err = ta.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

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

	err = ta.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}
