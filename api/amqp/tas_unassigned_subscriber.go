package amqp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
)

type TaskUnassignedEventSubscriber struct {
	subscriber message.Subscriber
}

func NewTaskUnassignedEventSubscriber(s message.Subscriber) (*TaskUnassignedEventSubscriber, error) {
	return &TaskUnassignedEventSubscriber{
		subscriber: s,
	}, nil
}

func (s TaskUnassignedEventSubscriber) Subscribe(ctx context.Context) (<-chan *api.TaskUnassignedEvent, error) {
	events := make(chan *api.TaskUnassignedEvent, 10) // let's buffer this channel for now...

	msgs, err := s.subscriber.Subscribe(ctx, api.TaskUnassignedEventName)
	if err != nil {
		return nil, err
	}

	// listen & map incoming messages to API contract
	go func() {
		defer close(events)

		for msg := range msgs {
			var event api.TaskUnassignedEvent
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
