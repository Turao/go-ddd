package ddd

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type DomainEvent interface {
	events.Event
	AggregateID() string
}

type DomainEventStore interface {
	Push(ctx context.Context, evt DomainEvent, expectedVersion int) error
	Events(ctx context.Context) ([]DomainEvent, error)
}
