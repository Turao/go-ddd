package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/events"
)

type RemoveTaskFromUserCommandHandler struct {
	repository account.Repository
	eventStore events.EventStore
}

var _ application.RemoveTaskFromUserCommandHandler = (*RemoveTaskFromUserCommandHandler)(nil)

func NewRemoveTaskFromUserCommandHandler(repository account.Repository, es events.EventStore) *RemoveTaskFromUserCommandHandler {
	return &RemoveTaskFromUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h RemoveTaskFromUserCommandHandler) Handle(ctx context.Context, req application.RemoveTaskFromUserCommand) error {
	events, err := h.eventStore.EventsByAggregateID(ctx, req.UserID)
	if err != nil {
		return err
	}

	aa, err := account.NewAccountAggregate(nil, h.eventStore)
	if err != nil {
		return nil
	}

	for _, event := range events {
		err = aa.HandleEvent(event)
		if err != nil {
			return err
		}
	}

	err = aa.RemoveTask(req.TaskID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *aa.Account)
	if err != nil {
		return err
	}

	return nil
}
