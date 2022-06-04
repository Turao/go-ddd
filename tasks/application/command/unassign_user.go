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

type UnassignUserCommandHandler struct {
	repository                   task.Repository
	eventStore                   events.EventStore
	taskUnassignedEventPublisher api.TaskUnassignedEventPublisher
}

func NewUnassignUserCommandHandler(repository task.Repository, es events.EventStore, tuep api.TaskUnassignedEventPublisher) *UnassignUserCommandHandler {
	return &UnassignUserCommandHandler{
		repository:                   repository,
		eventStore:                   es,
		taskUnassignedEventPublisher: tuep,
	}
}

func (h UnassignUserCommandHandler) Handle(ctx context.Context, req application.UnassignUserCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	// get assigned user's id so we can publish the integration event later
	assignedUser := *t.AssignedUser

	agg := task.NewTaskAggregate(task.TaskEventFactory{})
	root, err := ddd.NewAggregateRoot(agg)
	if err != nil {
		return err
	}

	err = root.HandleCommand(ctx, task.UnassignCommand{})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *agg.Task)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskUnassignedEvent(uuid.NewString(), req.TaskID, assignedUser)
	if err != nil {
		return err
	}

	err = h.taskUnassignedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
