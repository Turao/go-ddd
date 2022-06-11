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
	agg, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.UpdateDescriptionCommand{
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	return nil
}
