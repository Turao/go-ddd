package query

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type GetAccountDetailsQueryHandler struct {
	repository ddd.Repository
	eventStore events.EventStore
}

var _ application.GetAccountDetailsQueryHandler = (*GetAccountDetailsQueryHandler)(nil)

func NewGetAccountDetailsQueryHandler(repository ddd.Repository, es events.EventStore) *GetAccountDetailsQueryHandler {
	return &GetAccountDetailsQueryHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h GetAccountDetailsQueryHandler) Handle(ctx context.Context, req application.GetAccountDetailsQuery) (*application.GetAccountDetailsResponse, error) {
	agg, err := h.repository.FindByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &application.GetAccountDetailsResponse{
		Account: application.Account{
			ID: agg.ID(),
		},
	}, nil
}
