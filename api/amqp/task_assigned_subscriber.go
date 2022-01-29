package amqp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
)

type TaskAssignedEventSubscriber struct {
	subscriber message.Subscriber
}

func NewTaskAssignedEventSubscriber(s message.Subscriber) (*TaskAssignedEventSubscriber, error) {
	return &TaskAssignedEventSubscriber{
		subscriber: s,
	}, nil
}

func (s TaskAssignedEventSubscriber) Subscribe(ctx context.Context) (<-chan *api.TaskAssignedEvent, error) {
	events := make(chan *api.TaskAssignedEvent, 10) // let's buffer this channel for now...

	msgs, err := s.subscriber.Subscribe(ctx, api.TaskAssignedEventName)
	if err != nil {
		return nil, err
	}

	// listen & map incoming messages to API contract
	go func() {
		defer close(events)

		for msg := range msgs {
			var event api.TaskAssignedEvent
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
