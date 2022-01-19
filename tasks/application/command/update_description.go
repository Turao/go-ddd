package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateDescriptionCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewUpdateDescriptionCommandHandler(repository task.Repository, es events.EventStore) *UpdateDescriptionCommandHandler {
	return &UpdateDescriptionCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h UpdateDescriptionCommandHandler) Handle(ctx context.Context, req application.UpdateDescriptionCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	ta, err := task.NewTaskAggregate(t, h.eventStore)
	if err != nil {
		return err
	}

	err = ta.UpdateDescription(req.Description)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ta.Task)
	if err != nil {
		return err
	}

	return nil
}
