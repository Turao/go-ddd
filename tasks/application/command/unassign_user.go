package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	task "github.com/turao/go-ddd/tasks/domain"
)

type UnassignUserCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewUnassignUserCommandHandler(repository task.Repository, es events.EventStore) *UnassignUserCommandHandler {
	return &UnassignUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h UnassignUserCommandHandler) Handle(ctx context.Context, req application.UnassignUserCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	ta, err := task.NewTaskAggregate(t, h.eventStore)
	if err != nil {
		return err
	}

	err = ta.Unassign()
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ta.Task)
	if err != nil {
		return err
	}

	return nil
}
