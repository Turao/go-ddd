package ddd

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type DomainEvent interface {
	events.Event
	AggregateID() string
}

type Aggregate interface {
	ID() string
	Version() string
}

type AggregateRoot interface {
	ID() string
	Version() int
	HandleEvent(ctx context.Context, evt DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) ([]DomainEvent, error)
	ReplayEvents(ctx context.Context) error
	CommitEvents()
}
