package kafka

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
)

type TaskAssignedEventPublisher struct {
	publisher message.Publisher
}

func NewTaskAssignedEventPublisher(p message.Publisher) (*TaskAssignedEventPublisher, error) {
	return &TaskAssignedEventPublisher{
		publisher: p,
	}, nil
}

func (p TaskAssignedEventPublisher) Publish(ctx context.Context, event api.TaskAssignedEvent) error {
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
