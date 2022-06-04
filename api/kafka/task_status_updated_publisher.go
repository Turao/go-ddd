package kafka

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type TaskStatusUpdatedEventPublisher struct {
	publisher message.Publisher
}

func NewTaskStatusUpdatedEventPublisher(p message.Publisher) (*TaskStatusUpdatedEventPublisher, error) {
	return &TaskStatusUpdatedEventPublisher{
		publisher: p,
	}, nil
}

func (p TaskStatusUpdatedEventPublisher) Publish(ctx context.Context, event api.TaskStatusUpdatedEvent) error {
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
