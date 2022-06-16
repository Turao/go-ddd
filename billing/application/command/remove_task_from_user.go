package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
)

type RemoveTaskFromUserCommandHandler struct {
	repository ddd.Repository
	eventStore ddd.DomainEventStore
}

var _ application.RemoveTaskFromUserCommandHandler = (*RemoveTaskFromUserCommandHandler)(nil)

func NewRemoveTaskFromUserCommandHandler(repository ddd.Repository, es ddd.DomainEventStore) *RemoveTaskFromUserCommandHandler {
	return &RemoveTaskFromUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h RemoveTaskFromUserCommandHandler) Handle(ctx context.Context, req application.RemoveTaskFromUserCommand) error {
	agg, err := h.repository.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, account.RemoveTaskCommand{
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
