package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectHandler struct {
	eventStore events.EventStore
}

func NewDeleteProjectCommandHandler(es events.EventStore) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		eventStore: es,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req application.DeleteProjectCommand) error {
	evt, err := project.NewProjectDeletedEvent(req.ID)
	if err != nil {
		return err
	}

	return h.eventStore.Push(context.Background(), *evt)
}
