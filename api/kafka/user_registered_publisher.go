package kafka

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type UserRegisteredEventPublisher struct {
	publisher message.Publisher
}

func NewUserRegisteredEventPublisher(p message.Publisher) (*UserRegisteredEventPublisher, error) {
	return &UserRegisteredEventPublisher{
		publisher: p,
	}, nil
}

func (p UserRegisteredEventPublisher) Publish(ctx context.Context, event api.UserRegisteredEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		event.Name,
		message.NewMessage(uuid.NewString(), payload),
	)
	if err != nil {
		return err
	}

	return nil
}
