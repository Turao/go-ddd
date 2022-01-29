package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/user"
	"github.com/turao/go-ddd/events"
)

type RegisterUserCommandHandler struct {
	repository user.Repository
	eventStore events.EventStore
}

var _ application.RegisterUserCommandHandler = (*RegisterUserCommandHandler)(nil)

func NewRegisterUserCommandHandler(repository user.Repository, es events.EventStore) (*RegisterUserCommandHandler, error) {
	return &RegisterUserCommandHandler{
		repository: repository,
		eventStore: es,
	}, nil
}

func (h RegisterUserCommandHandler) Handle(ctx context.Context, req application.RegisterUserCommand) error {
	ua, err := user.NewUserAggregate(nil, h.eventStore)
	if err != nil {
		return nil
	}

	err = ua.RegisterUser(req.UserID)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ua.User)
	if err != nil {
		return err
	}

	return nil
}
