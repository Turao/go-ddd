package amqp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
)

type TaskStatusUpdatedEventSubscriber struct {
	subscriber message.Subscriber
}

func NewTaskStatusUpdatedEventSubscriber(s message.Subscriber) (*TaskStatusUpdatedEventSubscriber, error) {
	return &TaskStatusUpdatedEventSubscriber{
		subscriber: s,
	}, nil
}

func (s TaskStatusUpdatedEventSubscriber) Subscribe(ctx context.Context) (<-chan *api.TaskStatusUpdatedEvent, error) {
	events := make(chan *api.TaskStatusUpdatedEvent, 10) // let's buffer this channel for now...

	msgs, err := s.subscriber.Subscribe(ctx, api.TaskStatusUpdatedEventName)
	if err != nil {
		return nil, err
	}

	// listen & map incoming messages to API contract
	go func() {
		defer close(events)

		for msg := range msgs {
			var event api.TaskStatusUpdatedEvent
			err = json.Unmarshal(msg.Payload, &event)
			if err != nil {
				log.Println("[todo: dep. inject] unparseable message: ", string(msg.Payload), err)
				continue
			}

			events <- &event
			msg.Ack()
		}
	}()

	return events, nil
}
