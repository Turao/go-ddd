package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectHandler struct {
	repository project.Repository
	eventStore events.EventStore
}

func NewCreateProjectCommandHandler(repository project.Repository, es events.EventStore) *CreateProjectHandler {
	return &CreateProjectHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h *CreateProjectHandler) Handle(ctx context.Context, req application.CreateProjectCommand) error {
	agg, err := project.NewProjectAggregate(project.ProjectEventFactory{})
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, project.CreateProjectCommand{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
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
