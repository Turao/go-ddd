package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type RemoveTaskFromUserCommandHandler struct {
	repository account.Repository
	eventStore events.EventStore
}

var _ application.RemoveTaskFromUserCommandHandler = (*RemoveTaskFromUserCommandHandler)(nil)

func NewRemoveTaskFromUserCommandHandler(repository account.Repository, es events.EventStore) *RemoveTaskFromUserCommandHandler {
	return &RemoveTaskFromUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h RemoveTaskFromUserCommandHandler) Handle(ctx context.Context, req application.RemoveTaskFromUserCommand) error {
	agg := account.NewAccountAggregate(account.AccountEventsFactory{})
	root, err := ddd.NewAggregateRoot(agg, h.eventStore)
	if err != nil {
		return nil
	}

	err = root.HandleCommand(ctx, account.RemoveTaskFromUserCommand{
		TaskID: req.TaskID,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *agg.Account)
	if err != nil {
		return err
	}

	return nil
}
