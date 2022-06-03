package amqp

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type AMQPTaskAssignedEventPublisher struct {
	publisher message.Publisher
}

func NewAMQPTaskAssignedEventPublisher(p message.Publisher) (*AMQPTaskAssignedEventPublisher, error) {
	return &AMQPTaskAssignedEventPublisher{
		publisher: p,
	}, nil
}

func (p AMQPTaskAssignedEventPublisher) Publish(ctx context.Context, event api.TaskAssignedEvent) error {
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
