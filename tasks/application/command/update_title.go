package command

import (
	"context"

	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateTitleCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewUpdateTitleCommandHandler(repository task.Repository, es events.EventStore) *UpdateTitleCommandHandler {
	return &UpdateTitleCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h UpdateTitleCommandHandler) Handle(ctx context.Context, req application.UpdateTitleCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	agg := task.NewTaskAggregate(task.TaskEventFactory{})
	root, err := ddd.NewAggregateRoot(agg, h.eventStore)
	if err != nil {
		return err
	}

	agg.Task = t // todo: fix

	err = root.HandleCommand(ctx, task.UpdateTitleCommand{
		Title: req.Title,
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
