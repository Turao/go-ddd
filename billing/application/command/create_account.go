package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/ddd/eventsource"
)

type CreateAccountCommandHandler struct {
	repository ddd.Repository
	eventStore ddd.DomainEventStore
}

var _ application.CreateAccountCommandHandler = (*CreateAccountCommandHandler)(nil)

func NewCreateAccountCommandHandler(repository ddd.Repository, es ddd.DomainEventStore) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h CreateAccountCommandHandler) Handle(ctx context.Context, req application.CreateAccountCommand) error {
	agg, err := account.NewAggregate(account.AccountEventsFactory{}, account.WithAggregateID(req.UserID))
	if err != nil {
		return err
	}

	root, err := eventsource.NewAggregate(agg, h.eventStore)
	if err != nil {
		return nil
	}

	_, err = root.HandleCommand(ctx, account.CreateAccountCommand{
		UserID: req.UserID,
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
