package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
)

type CreateTaskCommandHandler struct {
	eventStore events.EventStore
}

func NewCreateTaskCommandHandler(es events.EventStore) (*CreateTaskCommandHandler, error) {
	return &CreateTaskCommandHandler{
		eventStore: es,
	}, nil
}

func (cth *CreateTaskCommandHandler) Handle(ctx context.Context, req application.CreateTaskCommand) error {
	return nil
}
