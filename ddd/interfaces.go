package ddd

import (
	"context"
	"encoding/json"

	"github.com/turao/go-ddd/events"
)

type DomainEvent interface {
	events.Event
	AggregateID() string
}

type DomainEventStore interface {
	json.Marshaler
	Push(ctx context.Context, evt DomainEvent, expectedVersion int) error
	Events(ctx context.Context, aggregateID string) ([]DomainEvent, error)
}
