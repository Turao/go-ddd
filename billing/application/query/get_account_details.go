package query

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/events/ddd"
)

type GetAccountDetailsQueryHandler struct {
	eventStore events.EventStore
}

var _ application.GetAccountDetailsQueryHandler = (*GetAccountDetailsQueryHandler)(nil)

func NewGetAccountDetailsQueryHandler(es events.EventStore) *GetAccountDetailsQueryHandler {
	return &GetAccountDetailsQueryHandler{
		eventStore: es,
	}
}

func (h GetAccountDetailsQueryHandler) Handle(ctx context.Context, req application.GetAccountDetailsQuery) (*application.GetAccountDetailsResponse, error) {
	agg := account.NewAccountAggregate(account.AccountEventsFactory{})
	root, err := ddd.NewAggregateRoot(agg, h.eventStore)
	if err != nil {
		return nil, err
	}

	return &application.GetAccountDetailsResponse{
		Account: application.Account{
			ID: root.ID(),
		},
	}, nil
}
