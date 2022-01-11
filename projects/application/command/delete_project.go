package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectCommand struct {
	ID project.ProjectID `json:"id"`
}

type DeleteProjectHandler struct {
	repo       project.Repository
	eventStore events.EventStore
}

func NewDeleteProjectCommandHandler(repo project.Repository, es events.EventStore) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		repo:       repo,
		eventStore: es,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req DeleteProjectCommand) error {
	evt, err := project.NewProjectDeletedEvent(req.ID)
	if err != nil {
		return err
	}

	return h.eventStore.Push(context.Background(), evt)
}
