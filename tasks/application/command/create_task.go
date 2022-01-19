package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type CreateTaskCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewCreateTaskCommandHandler(repository task.Repository, es events.EventStore) *CreateTaskCommandHandler {
	return &CreateTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h *CreateTaskCommandHandler) Handle(ctx context.Context, req application.CreateTaskCommand) error {
	t, err := task.CreateTask(req.ProjectID, req.Title, req.Description)
	if err != nil {
		return err
	}

	evt, err := task.NewTaskCreatedEvent(t.ID, t.ProjectID, t.Title, t.Description)
	if err != nil {
		return err
	}

	err = h.eventStore.Push(ctx, *evt)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *t)
	if err != nil {
		return err
	}

	return nil
}
