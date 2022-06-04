package command

import (
	"context"

	"github.com/turao/go-ddd/ddd"
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

	agg := task.NewTaskAggregate(task.TaskEventFactory{})
	root, err := ddd.NewAggregateRoot(agg)
	if err != nil {
		return err
	}

	agg.Task = t // todo: fix

	err = root.HandleCommand(ctx, task.UpdateDescriptionCommand{
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *agg.Task)
	if err != nil {
		return err
	}

	return nil
}
