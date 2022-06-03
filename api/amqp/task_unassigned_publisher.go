package amqp

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type AMQPTaskUnassignedEventPublisher struct {
	publisher message.Publisher
}

func NewAMQPTaskUnassignedEventPublisher(p message.Publisher) (*AMQPTaskUnassignedEventPublisher, error) {
	return &AMQPTaskUnassignedEventPublisher{
		publisher: p,
	}, nil
}

func (p AMQPTaskUnassignedEventPublisher) Publish(ctx context.Context, event api.TaskUnassignedEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		event.Name(),
		message.NewMessage(uuid.NewString(), payload),
	)
	if err != nil {
		return err
	}

	return nil
}
