package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectHandler struct {
	eventStore events.EventStore
}

func NewCreateProjectCommandHandler(es events.EventStore) *CreateProjectHandler {
	return &CreateProjectHandler{
		eventStore: es,
	}
}

func (h *CreateProjectHandler) Handle(ctx context.Context, req application.CreateProjectCommand) error {
	pa, err := project.NewProjectAggregate(nil, h.eventStore)
	if err != nil {
		return err
	}

	err = pa.CreateProject(req.Name)
	if err != nil {
		return err
	}

	return nil
}
