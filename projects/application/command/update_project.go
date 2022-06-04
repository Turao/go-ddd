package command

import (
	"context"

	"github.com/turao/go-ddd/ddd"
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
	p, err := h.repository.FindProjectByID(ctx, req.ID)
	if err != nil {
		return err
	}

	pa := project.NewProjectAggregate(project.ProjectAggregate{})
	pa.Project = p // todo: load from snapshot instead of overriding directly

	root, err := ddd.NewAggregateRoot(pa, ddd.WithEventStore(h.eventStore))
	if err != nil {
		return err
	}

	err = root.HandleCommand(ctx, project.UpdateProjectCommand{
		Name: req.Name,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *pa.Project)
	if err != nil {
		return err
	}

	return nil
}
