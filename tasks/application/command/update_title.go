package command

import (
	"context"

	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateTitleCommandHandler struct {
	repository task.Repository
}

func NewUpdateTitleCommandHandler(repository task.Repository) *UpdateTitleCommandHandler {
	return &UpdateTitleCommandHandler{
		repository: repository,
	}
}

func (h UpdateTitleCommandHandler) Handle(ctx context.Context, req application.UpdateTitleCommand) error {
	agg, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.UpdateTitleCommand{
		Title: req.Title,
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
