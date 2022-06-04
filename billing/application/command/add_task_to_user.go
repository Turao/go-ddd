package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type AddTaskToUserCommandHandler struct {
	repository account.Repository
	eventStore events.EventStore
}

var _ application.AddTaskToUserCommandHandler = (*AddTaskToUserCommandHandler)(nil)

func NewAddTaskToUserCommandHandler(repository account.Repository, es events.EventStore) *AddTaskToUserCommandHandler {
	return &AddTaskToUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AddTaskToUserCommandHandler) Handle(ctx context.Context, req application.AddTaskToUserCommand) error {
	agg := account.NewAccountAggregate(account.AccountEventsFactory{})
	root, err := ddd.NewAggregateRoot(agg, ddd.WithEventStore(h.eventStore))
	if err != nil {
		return nil
	}

	err = root.HandleCommand(ctx, account.AddTaskToUserCommand{
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
