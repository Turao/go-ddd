package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
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
	aa, err := account.NewAccountAggregate(nil, h.eventStore)
	if err != nil {
		return nil
	}

	err = aa.CreateAccount(req.AccountID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *aa.Account)
	if err != nil {
		return err
	}

	return nil
}
