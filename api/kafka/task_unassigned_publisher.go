package kafka

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type TaskUnassignedEventPublisher struct {
	publisher message.Publisher
}

func NewTaskUnassignedEventPublisher(p message.Publisher) (*TaskUnassignedEventPublisher, error) {
	return &TaskUnassignedEventPublisher{
		publisher: p,
	}, nil
}

func (p TaskUnassignedEventPublisher) Publish(ctx context.Context, event api.TaskUnassignedEvent) error {
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
