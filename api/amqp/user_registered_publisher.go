package amqp

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type AMQPUserRegisteredEventPublisher struct {
	publisher message.Publisher
}

func NewAMQPUserRegisteredEventPublisher(p message.Publisher) (*AMQPUserRegisteredEventPublisher, error) {
	return &AMQPUserRegisteredEventPublisher{
		publisher: p,
	}, nil
}

func (p AMQPUserRegisteredEventPublisher) Publish(ctx context.Context, event api.UserRegisteredEvent) error {
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
