package amqp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/go-ddd/api"
)

type UserRegisteredEventSubscriber struct {
	subscriber message.Subscriber
}

func NewUserRegisteredEventSubscriber(s message.Subscriber) (*UserRegisteredEventSubscriber, error) {
	return &UserRegisteredEventSubscriber{
		subscriber: s,
	}, nil
}

func (s UserRegisteredEventSubscriber) Subscribe(ctx context.Context) (<-chan *api.UserRegisteredEvent, error) {
	events := make(chan *api.UserRegisteredEvent, 10) // let's buffer this channel for now...

	msgs, err := s.subscriber.Subscribe(ctx, api.UserRegisteredEventName)
	if err != nil {
		return nil, err
	}

	// listen & map incoming messages to API contract
	go func() {
		defer close(events)

		for msg := range msgs {
			var event api.UserRegisteredEvent
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
