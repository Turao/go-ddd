package amqp

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type AMQPTaskStatusUpdatedEventPublisher struct {
	publisher message.Publisher
}

func NewAMQPTaskStatusUpdatedEventPublisher(p message.Publisher) (*AMQPTaskStatusUpdatedEventPublisher, error) {
	return &AMQPTaskStatusUpdatedEventPublisher{
		publisher: p,
	}, nil
}

func (p AMQPTaskStatusUpdatedEventPublisher) Publish(ctx context.Context, event api.TaskStatusUpdatedEvent) error {
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
