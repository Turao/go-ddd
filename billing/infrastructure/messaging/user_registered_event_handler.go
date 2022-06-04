package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type UserRegisteredEventHandler struct {
	CommandHandler application.CreateAccountCommandHandler
}

func (h UserRegisteredEventHandler) Handle(msg *message.Message) error {
	var evt api.UserRegisteredEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.CreateAccountCommand{
		AccountID: evt.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}
