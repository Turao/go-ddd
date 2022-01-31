package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type TaskAssignedEventHandler struct {
	CommandHandler application.AddTaskToUserCommandHandler
}

func (h TaskAssignedEventHandler) Handle(msg *message.Message) error {
	var evt api.TaskAssignedEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.AddTaskToUserCommand{
		UserID: evt.UserID,
		TaskID: evt.TaskID,
	})

	if err != nil {
		return err
	}

	return nil
}
