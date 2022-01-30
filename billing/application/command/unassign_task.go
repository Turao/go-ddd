package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/user"
	"github.com/turao/go-ddd/events"
)

type UnassignTaskCommandHandler struct {
	repository user.Repository
	eventStore events.EventStore
}

var _ application.UnassignTaskCommandHandler = (*UnassignTaskCommandHandler)(nil)

func NewUnassignTaskCommandHandler(repository user.Repository, es events.EventStore) *UnassignTaskCommandHandler {
	return &UnassignTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h UnassignTaskCommandHandler) Handle(ctx context.Context, req application.UnassignTaskCommand) error {
	u, err := h.repository.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	ua, err := user.NewUserAggregate(u, h.eventStore)
	if err != nil {
		return nil
	}

	err = ua.UnassignTask(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ua.User)
	if err != nil {
		return err
	}

	return nil
}
