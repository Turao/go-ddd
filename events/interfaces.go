package events

import (
	"context"
	"time"
)

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccuredAt() time.Time
}

// ----

type AggregateID = string

type DomainEvent interface {
	Event
	AggregateID() AggregateID
}

// ----

type CorrelationID = string

type IntegrationEvent interface {
	Event
	CorrelationID() CorrelationID
}

// ---
type EventStore interface {
	Push(ctx context.Context, evt Event) error
	Events(ctx context.Context) ([]Event, error)
	EventsByAggregateID(ctx context.Context, aggregateID AggregateID) ([]DomainEvent, error)
}
