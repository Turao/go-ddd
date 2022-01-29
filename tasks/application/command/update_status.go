package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateStatusCommandHandler struct {
	repository                      task.Repository
	eventStore                      events.EventStore
	taskStatusUpdatedEventPublisher api.TaskStatusUpdatedEventPublisher
}

func NewUpdateStatusCommandHandler(repository task.Repository, es events.EventStore, tsuep api.TaskStatusUpdatedEventPublisher) *UpdateStatusCommandHandler {
	return &UpdateStatusCommandHandler{
		repository:                      repository,
		eventStore:                      es,
		taskStatusUpdatedEventPublisher: tsuep,
	}
}

func (h UpdateStatusCommandHandler) Handle(ctx context.Context, req application.UpdateStatusCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	ta, err := task.NewTaskAggregate(t, h.eventStore)
	if err != nil {
		return err
	}

	err = ta.UpdateStatus(req.Status)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ta.Task)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskStatusUpdatedEvent(uuid.NewString(), t.ID, t.Status)
	if err != nil {
		return err
	}

	err = h.taskStatusUpdatedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
