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
	p, err := project.CreateProject(req.Name)
	if err != nil {
		return err
	}

	evt, err := project.NewProjectCreatedEvent(p.ID, p.Name)
	if err != nil {
		return err
	}

	return h.eventStore.Push(context.Background(), *evt)
}
