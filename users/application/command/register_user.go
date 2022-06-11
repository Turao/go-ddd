package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/domain/user"
)

type RegisterUserHandler struct {
	repository                   user.Repository
	eventStore                   events.EventStore
	userRegisteredEventPublisher api.UserRegisteredEventPublisher
}

func NewRegisterUserHandler(
	repository user.Repository,
	es events.EventStore,
	urep api.UserRegisteredEventPublisher,
) *RegisterUserHandler {
	return &RegisterUserHandler{
		repository:                   repository,
		eventStore:                   es,
		userRegisteredEventPublisher: urep,
	}
}

func (h RegisterUserHandler) Handle(ctx context.Context, req application.RegisterUserCommand) error {
	agg := user.NewUserAggregate(user.UserEventsFactory{})
	root, err := ddd.NewAggregateRoot(
		agg,
		h.eventStore,
	)
	if err != nil {
		return err
	}

	_, err = root.HandleCommand(ctx, user.RegisterUserCommand{
		Username: req.Username,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	ie, err := api.NewUserRegisteredEvent(uuid.NewString(), agg.ID())
	if err != nil {
		return err
	}

	err = h.userRegisteredEventPublisher.Publish(
		ctx,
		*ie,
	)
	if err != nil {
		return err
	}

	return nil
}
