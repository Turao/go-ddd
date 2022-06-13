package events

import (
	"context"
	"time"
)

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccurredAt() time.Time
}

// type IntegrationEvent interface {
// 	Event
// 	CorrelationID() string
// }

type EventStore interface {
	Push(ctx context.Context, evt Event, expectedVersion int) error
	Events(ctx context.Context) ([]Event, error)
}
