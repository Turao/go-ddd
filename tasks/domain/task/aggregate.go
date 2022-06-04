package task

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type TaskAggregate struct {
	Task *Task
	EventFactory
}

var (
	ErrUnknownEvent   = errors.New("unknown event")
	ErrUnknownCommand = errors.New("unknown command")
)

func NewTaskAggregate(ef EventFactory) *TaskAggregate {
	return &TaskAggregate{
		Task:         nil,
		EventFactory: ef,
	}
}

func (ta *TaskAggregate) ID() string {
	return ta.Task.ID
}

func (ta *TaskAggregate) HandleEvent(ctx context.Context, e ddd.DomainEvent) error {
	switch event := e.(type) {
	case TaskCreatedEvent:
		t, err := NewTask(event.AggregateID(), event.ProjectID, event.Title, event.Description)
		if err != nil {
			return err
		}
		ta.Task = t
		return nil
	case TaskAssignedEvent:
		err := ta.Task.AssignToUser(event.AssignedTo)
		if err != nil {
			return err
		}
		return nil
	case TaskUnassignedEvent:
		err := ta.Task.Unassign()
		if err != nil {
			return err
		}
		return nil
	case TitleUpdatedEvent:
		err := ta.Task.UpdateTitle(event.Title)
		if err != nil {
			return err
		}
		return nil
	case DescriptionUpdatedEvent:
		err := ta.Task.UpdateDescription(event.Description)
		if err != nil {
			return err
		}
		return nil
	case StatusUpdatedEvent:
		err := ta.Task.UpdateStatus(event.Status)
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrUnknownEvent
	}
}

func (ta *TaskAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case CreateTaskCommand:
		return ta.handleCreateTaskCommand(ctx, c)
	case AssignToUserCommand:
		return ta.handleAssignToUserCommand(ctx, c)
	case UnassignCommand:
		return ta.handleUnassignCommand(ctx, c)
	case UpdateTitleCommand:
		return ta.handleUpdateTitleCommand(ctx, c)
	case UpdateDescriptionCommand:
		return ta.handleUpdateDescriptionCommand(ctx, c)
	case UpdateStatusCommand:
		return ta.handleUpdateStatusCommand(ctx, c)
	default:
		return nil, ErrUnknownCommand
	}
}

func (ta *TaskAggregate) handleCreateTaskCommand(ctx context.Context, cmd CreateTaskCommand) ([]ddd.DomainEvent, error) {
	t, err := NewTask(uuid.NewString(), cmd.ProjectID, cmd.Title, cmd.Description)
	if err != nil {
		return nil, err
	}

	ta.Task = t

	evt, err := ta.EventFactory.NewTaskCreatedEvent(t.ID, t.ProjectID, t.Title, t.Description)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) handleAssignToUserCommand(ctx context.Context, cmd AssignToUserCommand) ([]ddd.DomainEvent, error) {
	err := ta.Task.AssignToUser(cmd.UserID)
	if err != nil {
		return nil, err
	}

	evt, err := ta.EventFactory.NewTaskAssignedEvent(ta.Task.ID, cmd.UserID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) handleUnassignCommand(ctx context.Context, cmd UnassignCommand) ([]ddd.DomainEvent, error) {
	err := ta.Task.Unassign()
	if err != nil {
		return nil, err
	}

	evt, err := ta.EventFactory.NewTaskUnassignedEvent(ta.Task.ID)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) handleUpdateTitleCommand(ctx context.Context, cmd UpdateTitleCommand) ([]ddd.DomainEvent, error) {
	err := ta.Task.UpdateTitle(cmd.Title)
	if err != nil {
		return nil, err
	}

	evt, err := ta.EventFactory.NewTitleUpdatedEvent(ta.Task.ID, cmd.Title)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) handleUpdateDescriptionCommand(ctx context.Context, cmd UpdateDescriptionCommand) ([]ddd.DomainEvent, error) {
	err := ta.Task.UpdateDescription(cmd.Description)
	if err != nil {
		return nil, err
	}

	evt, err := ta.EventFactory.NewDescriptionUpdatedEvent(ta.Task.ID, cmd.Description)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) handleUpdateStatusCommand(ctx context.Context, cmd UpdateStatusCommand) ([]ddd.DomainEvent, error) {
	err := ta.Task.UpdateStatus(cmd.Status)
	if err != nil {
		return nil, err
	}

	evt, err := ta.EventFactory.NewStatusUpdatedEvent(ta.Task.ID, cmd.Status)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}

func (ta TaskAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Task Task `json:"task"`
	}{
		Task: *ta.Task,
	})
}

func (ta *TaskAggregate) UnmarshalJSON(data []byte) error {
	var payload struct {
		Task Task `json:"Task"`
	}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}
	ta.Task = &payload.Task
	return nil
}
