package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/domain/user"
)

type RegisterUserHandler struct {
	eventStore events.EventStore
}

func NewRegisterUserHandler(es events.EventStore) (*RegisterUserHandler, error) {
	return &RegisterUserHandler{
		eventStore: es,
	}, nil
}

func (h RegisterUserHandler) Handle(ctx context.Context, req application.RegisterUserCommand) error {
	ua, err := user.NewUserAggregate(nil, h.eventStore)
	if err != nil {
		return err
	}

	err = ua.RegisterUser(req.Username)
	if err != nil {
		return err
	}

	return nil
}
