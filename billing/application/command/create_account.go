package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type CreateAccountCommandHandler struct {
	repository account.Repository
	eventStore events.EventStore
}

var _ application.CreateAccountCommandHandler = (*CreateAccountCommandHandler)(nil)

func NewCreateAccountCommandHandler(repository account.Repository, es events.EventStore) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h CreateAccountCommandHandler) Handle(ctx context.Context, req application.CreateAccountCommand) error {
	agg := account.NewAccountAggregate(account.AccountEventsFactory{})
	root, err := ddd.NewAggregateRoot(agg, ddd.WithEventStore(h.eventStore))
	if err != nil {
		return nil
	}

	err = root.HandleCommand(ctx, account.CreateAccountCommand{
		UserID: req.AccountID,
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
