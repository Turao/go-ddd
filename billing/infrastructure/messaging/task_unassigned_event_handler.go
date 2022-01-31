package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type TaskUnassignedEventHandler struct {
	CommandHandler application.RemoveTaskFromUserCommandHandler
}

func (h TaskUnassignedEventHandler) Handle(msg *message.Message) error {
	var evt api.TaskUnassignedEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.RemoveTaskFromUserCommand{
		UserID: evt.UserID,
		TaskID: evt.TaskID,
	})

	if err != nil {
		return err
	}

	return nil
}
