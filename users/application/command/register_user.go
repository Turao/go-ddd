package command

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/domain/user"
)

type RegisterUserHandler struct {
	repository     user.Repository
	eventStore     events.EventStore
	eventPublisher message.Publisher
}

func NewRegisterUserHandler(repository user.Repository, es events.EventStore, ep message.Publisher) *RegisterUserHandler {
	return &RegisterUserHandler{
		repository:     repository,
		eventStore:     es,
		eventPublisher: ep,
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

	payload, err := json.Marshal(ie)
	if err != nil {
		return err
	}

	err = h.eventPublisher.Publish(
		ie.Name(),
		message.NewMessage(uuid.NewString(), payload),
	)
	if err != nil {
		return err
	}

	return nil
}
