package messaging

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/users/application"
)

type OnRegisterUser struct {
	CommandHandler application.RegisterUserCommandHandler
}

type MockUserRegisteredEvent struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

func (r OnRegisterUser) Handle(msg *message.Message) ([]*message.Message, error) {
	var req application.RegisterUserCommand
	err := json.Unmarshal(msg.Payload, &req)
	if err != nil {
		return make([]*message.Message, 0), err
	}

	err = r.CommandHandler.Handle(context.Background(), req)
	if err != nil {
		return make([]*message.Message, 0), err
	}

	evt, err := json.Marshal(MockUserRegisteredEvent{
		UserID:   "todo-find-a-way-to-get-the-user-id",
		Username: req.Username,
	})
	if err != nil {
		return make([]*message.Message, 0), err
	}

	return []*message.Message{
		message.NewMessage(uuid.NewString(), evt),
	}, nil
}
