package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
)

type AddTaskToUserCommandHandler struct {
	repository ddd.Repository
	eventStore ddd.DomainEventStore
}

var _ application.AddTaskToUserCommandHandler = (*AddTaskToUserCommandHandler)(nil)

func NewAddTaskToUserCommandHandler(repository ddd.Repository, es ddd.DomainEventStore) *AddTaskToUserCommandHandler {
	return &AddTaskToUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AddTaskToUserCommandHandler) Handle(ctx context.Context, req application.AddTaskToUserCommand) error {
	agg, err := h.repository.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, account.AddTaskToUserCommand{
		TaskID: req.TaskID,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	return nil
}
