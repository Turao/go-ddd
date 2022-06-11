package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
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
	agg, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.UnassignCommand{})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskUnassignedEvent(uuid.NewString(), req.TaskID, *agg.Task.AssignedUser)
	if err != nil {
		return err
	}

	err = h.taskUnassignedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
