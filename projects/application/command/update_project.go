package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type UpdateProjectCommand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateProjectHandler struct {
	repo       project.Repository
	eventStore events.EventStore
}

func NewUpdateProjectCommandHandler(repo project.Repository, es events.EventStore) *UpdateProjectHandler {
	return &UpdateProjectHandler{
		repo:       repo,
		eventStore: es,
	}
}

func (h *UpdateProjectHandler) Handle(ctx context.Context, req UpdateProjectCommand) error {
	evt, err := project.NewProjectUpdatedEvent(req.ID, req.Name)
	if err != nil {
		return err
	}

	return h.eventStore.Push(context.Background(), *evt)
}
