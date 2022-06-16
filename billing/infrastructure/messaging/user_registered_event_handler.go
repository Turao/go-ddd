package messaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type UserRegisteredEventHandler struct {
	CommandHandler application.CreateAccountCommandHandler
}

func (h UserRegisteredEventHandler) Handle(msg *message.Message) error {
	var evt api.UserRegisteredEvent
	log.Println("msg.payload:", msg.Payload)
	log.Println("before:", evt)

	err := json.Unmarshal(msg.Payload, &evt)
	if err != nil {
		return err
	}

	log.Println("after:", evt)

	err = h.CommandHandler.Handle(context.Background(), application.CreateAccountCommand{
		UserID: evt.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}
