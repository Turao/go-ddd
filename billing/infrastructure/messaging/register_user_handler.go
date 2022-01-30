package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type RegisterUserCommandHandler struct {
	CommandHandler application.RegisterUserCommandHandler
}

func (h RegisterUserCommandHandler) Handle(msg *message.Message) error {
	var evt api.UserRegisteredEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.RegisterUserCommand{
		UserID: evt.AggregateID,
	})

	if err != nil {
		return err
	}

	return nil
}
