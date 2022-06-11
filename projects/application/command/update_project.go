package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type UpdateProjectHandler struct {
	repository project.Repository
	eventStore events.EventStore
}

func NewUpdateProjectCommandHandler(repository project.Repository, es events.EventStore) *UpdateProjectHandler {
	return &UpdateProjectHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h *UpdateProjectHandler) Handle(ctx context.Context, req application.UpdateProjectCommand) error {
	agg, err := h.repository.FindProjectByID(ctx, req.ID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, project.UpdateProjectCommand{
		Name: req.Name,
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
