package command

import (
	"context"

	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type CreateTaskCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewCreateTaskCommandHandler(repository task.Repository, es events.EventStore) *CreateTaskCommandHandler {
	return &CreateTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h *CreateTaskCommandHandler) Handle(ctx context.Context, req application.CreateTaskCommand) error {
	agg := task.NewTaskAggregate(task.TaskEventFactory{})
	root, err := ddd.NewAggregateRoot(agg)
	if err != nil {
		return err
	}

	err = root.HandleCommand(ctx, task.CreateTaskCommand{
		ProjectID:   req.ProjectID,
		Title:       req.Title,
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
