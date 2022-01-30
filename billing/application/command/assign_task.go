package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/user"
	"github.com/turao/go-ddd/events"
)

type AssignTaskCommandHandler struct {
	repository user.Repository
	eventStore events.EventStore
}

var _ application.AssignTaskCommandHandler = (*AssignTaskCommandHandler)(nil)

func NewAssignTaskCommandHandler(repository user.Repository, es events.EventStore) *AssignTaskCommandHandler {
	return &AssignTaskCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AssignTaskCommandHandler) Handle(ctx context.Context, req application.AssignTaskCommand) error {
	u, err := h.repository.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	ua, err := user.NewUserAggregate(u, h.eventStore)
	if err != nil {
		return nil
	}

	err = ua.AssignTask(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ua.User)
	if err != nil {
		return err
	}

	return nil
}
