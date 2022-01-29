package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type AssignToUserCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewAssignToUserCommandHandler(repository task.Repository, es events.EventStore) *AssignToUserCommandHandler {
	return &AssignToUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AssignToUserCommandHandler) Handle(ctx context.Context, req application.AssignToUserCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	ta, err := task.NewTaskAggregate(t, h.eventStore)
	if err != nil {
		return err
	}

	err = ta.AssignTo(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ta.Task)
	if err != nil {
		return err
	}

	return nil
}
