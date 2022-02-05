package ddd

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type AggregateRoot struct {
	Version int
	Events  events.EventStore
}

type EventHandlerFunc = func(event events.DomainEvent) error

func (ar *AggregateRoot) HandleEvent(event events.DomainEvent, fn EventHandlerFunc) error {
	err := fn(event)
	if err != nil {
		return err
	}

	ar.Version = ar.Version + 1
	return nil
}

func (ar *AggregateRoot) AddEvent(event events.DomainEvent) error {
	err := ar.Events.Push(context.Background(), event)
	if err != nil {
		return err
	}

	ar.Version = ar.Version + 1
	return nil
}
