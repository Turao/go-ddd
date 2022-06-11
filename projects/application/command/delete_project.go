package command

import (
	"context"

	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectHandler struct {
	repository project.Repository
}

func NewDeleteProjectCommandHandler(repository project.Repository) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		repository: repository,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req application.DeleteProjectCommand) error {
	agg, err := h.repository.FindProjectByID(ctx, req.ID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, project.DeleteProjectCommand{})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	return nil
}
