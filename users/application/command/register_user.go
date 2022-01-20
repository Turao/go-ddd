package command

import (
	"context"
	"log"

	"github.com/turao/go-ddd/api"
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

	ie, err := api.NewUserRegisteredEvent(ua.User.ID)
	if err != nil {
		return err
	}
	log.Println("TODO: publish this user.registered integration event: ", ie)

	return nil
}
