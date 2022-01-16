package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectHandler struct {
	repository project.Repository
	eventStore events.EventStore
}

func NewDeleteProjectCommandHandler(repository project.Repository, es events.EventStore) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req application.DeleteProjectCommand) error {
	p, err := h.repository.FindProjectByID(ctx, req.ID)
	if err != nil {
		return err
	}

	pa, err := project.NewProjectAggregate(p, h.eventStore)
	if err != nil {
		return err
	}

	err = pa.DeleteProject()
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *pa.Project)
	if err != nil {
		return err
	}

	return nil
}
