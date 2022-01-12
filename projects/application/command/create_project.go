package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectCommand struct {
	Name string `json:"name"`
}

type CreateProjectHandler struct {
	eventStore events.EventStore
}

func NewCreateProjectCommandHandler(es events.EventStore) *CreateProjectHandler {
	return &CreateProjectHandler{
		eventStore: es,
	}
}

func (h *CreateProjectHandler) Handle(ctx context.Context, req CreateProjectCommand) error {
	evt, err := project.NewProjectCreatedEvent(uuid.NewString(), req.Name)
	if err != nil {
		return err
	}

	return h.eventStore.Push(context.Background(), *evt)
}
