package command

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type CreateTaskCommandHandler struct {
	eventStore events.EventStore
}

func NewCreateTaskCommandHandler(es events.EventStore) *CreateTaskCommandHandler {
	return &CreateTaskCommandHandler{
		eventStore: es,
	}
}

func (cth *CreateTaskCommandHandler) Handle(ctx context.Context, req application.CreateTaskCommand) error {
	log.Println("creating task", req)
	evt, err := task.NewTaskCreatedEvent(uuid.NewString(), req.ID, req.Title, req.Description)
	if err != nil {
		return err
	}

	return cth.eventStore.Push(ctx, evt)
}
