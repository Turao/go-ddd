package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateStatusCommandHandler struct {
	repository                      task.Repository
	taskStatusUpdatedEventPublisher api.TaskStatusUpdatedEventPublisher
}

func NewUpdateStatusCommandHandler(repository task.Repository, tsuep api.TaskStatusUpdatedEventPublisher) *UpdateStatusCommandHandler {
	return &UpdateStatusCommandHandler{
		repository:                      repository,
		taskStatusUpdatedEventPublisher: tsuep,
	}
}

func (h UpdateStatusCommandHandler) Handle(ctx context.Context, req application.UpdateStatusCommand) error {
	agg, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.UpdateStatusCommand{
		Status: req.Status,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskStatusUpdatedEvent(uuid.NewString(), agg.ID(), agg.Task.Status)
	if err != nil {
		return err
	}

	err = h.taskStatusUpdatedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
