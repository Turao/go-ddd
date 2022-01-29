package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	task "github.com/turao/go-ddd/tasks/domain"
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
	ta, err := task.NewTaskAggregate(nil, h.eventStore)
	if err != nil {
		return err
	}

	t, err := ta.CreateTask(req.ProjectID, req.Title, req.Description)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *t)
	if err != nil {
		return err
	}

	return nil
}
