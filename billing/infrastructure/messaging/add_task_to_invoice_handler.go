package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type AddTaskToInvoiceCommandHandler struct {
	CommandHandler application.AddTaskToInvoiceCommandHandler
}

func (h AddTaskToInvoiceCommandHandler) Handle(msg *message.Message) error {
	var evt api.UserRegisteredEvent
	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	err = h.CommandHandler.Handle(context.Background(), application.AddTaskToInvoiceCommand{
		UserID: evt.AggregateID,
	})

	if err != nil {
		return err
	}

	return nil
}