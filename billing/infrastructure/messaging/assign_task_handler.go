package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type AssignTaskCommandHandler struct {
	CommandHandler application.AssignTaskCommandHandler
}

func (h AssignTaskCommandHandler) Handle(msg *message.Message) error {
	var evt api.UserRegisteredEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.AssignTaskCommand{
		UserID: evt.AggregateID,
	})

	if err != nil {
		return err
	}

	return nil
}
