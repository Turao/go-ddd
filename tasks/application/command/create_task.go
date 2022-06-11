package command

import (
	"context"

	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type CreateTaskCommandHandler struct {
	repository task.Repository
}

func NewCreateTaskCommandHandler(repository task.Repository) *CreateTaskCommandHandler {
	return &CreateTaskCommandHandler{
		repository: repository,
	}
}

func (h *CreateTaskCommandHandler) Handle(ctx context.Context, req application.CreateTaskCommand) error {
	agg, err := task.NewTaskAggregate(task.TaskEventFactory{})
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.CreateTaskCommand{
		ProjectID:   req.ProjectID,
		Title:       req.Title,
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
