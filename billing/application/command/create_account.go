package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/ddd/eventsource"
	"github.com/turao/go-ddd/events"
)

type CreateAccountCommandHandler struct {
	repository ddd.Repository
	eventStore events.EventStore
}

var _ application.CreateAccountCommandHandler = (*CreateAccountCommandHandler)(nil)

func NewCreateAccountCommandHandler(repository ddd.Repository, es events.EventStore) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h CreateAccountCommandHandler) Handle(ctx context.Context, req application.CreateAccountCommand) error {
	agg := account.NewAccountAggregate(account.AccountEventsFactory{})
	root, err := eventsource.NewAggregate(agg, h.eventStore)
	if err != nil {
		return nil
	}

	_, err = root.HandleCommand(ctx, account.CreateAccountCommand{
		UserID: req.AccountID,
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
