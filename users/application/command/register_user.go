package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/domain/user"
)

type RegisterUserHandler struct {
	repository user.Repository
	eventStore events.EventStore
}

func NewRegisterUserHandler(repository user.Repository, es events.EventStore) *RegisterUserHandler {
	return &RegisterUserHandler{
		repository: repository,
		eventStore: es,
	}
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

	err = h.repository.Save(ctx, *ua.User)
	if err != nil {
		return err
	}

	return nil
}
