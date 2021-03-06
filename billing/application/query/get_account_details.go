package query

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/ddd"
)

type GetAccountDetailsQueryHandler struct {
	repository ddd.Repository
	eventStore ddd.DomainEventStore
}

var _ application.GetAccountDetailsQueryHandler = (*GetAccountDetailsQueryHandler)(nil)

func NewGetAccountDetailsQueryHandler(repository ddd.Repository, es ddd.DomainEventStore) *GetAccountDetailsQueryHandler {
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
