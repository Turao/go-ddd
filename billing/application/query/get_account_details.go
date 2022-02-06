package query

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/events"
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
	events, err := h.eventStore.EventsByAggregateID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	aa, err := account.NewAccountAggregate(nil, h.eventStore)
	if err != nil {
		return nil, err
	}

	// replay events
	for _, event := range events {
		if err := aa.HandleEvent(event); err != nil {
			return nil, err
		}
	}

	return &application.GetAccountDetailsResponse{
		Account: application.Account{
			ID: req.AccountID,
		},
	}, nil
}
