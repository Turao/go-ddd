package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/ddd/eventsource"
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
	agg, err := account.NewAggregate(account.AccountEventsFactory{}, account.WithAggregateID(req.UserID))
	if err != nil {
		return err
	}

	root, err := eventsource.NewAggregate(agg, h.eventStore)
	if err != nil {
		return nil
	}

	err = root.ReplayEvents()
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, account.RemoveTaskCommand{
		TaskID: req.TaskID,
	})
	if err != nil {
		return err
	}

	err = root.CommitEvents()
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	return nil
}
