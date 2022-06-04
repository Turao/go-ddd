package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type AssignToUserCommandHandler struct {
	repository                 task.Repository
	eventStore                 events.EventStore
	taskAssignedEventPublisher api.TaskAssignedEventPublisher
}

func NewAssignToUserCommandHandler(repository task.Repository, es events.EventStore, taep api.TaskAssignedEventPublisher) *AssignToUserCommandHandler {
	return &AssignToUserCommandHandler{
		repository:                 repository,
		eventStore:                 es,
		taskAssignedEventPublisher: taep,
	}
}

func (h AssignToUserCommandHandler) Handle(ctx context.Context, req application.AssignToUserCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	agg := task.NewTaskAggregate(task.TaskEventFactory{})
	root, err := ddd.NewAggregateRoot(agg, h.eventStore)
	if err != nil {
		return err
	}

	agg.Task = t // todo: find a better way to load the entity

	err = root.HandleCommand(ctx, task.AssignToUserCommand{
		UserID: req.UserID,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *agg.Task)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskAssignedEvent(uuid.NewString(), req.TaskID, req.UserID)
	if err != nil {
		return err
	}

	err = h.taskAssignedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
