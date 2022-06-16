package ddd

import "context"

type DomainEventStore interface {
	Push(ctx context.Context, evt DomainEvent) error
	Events(ctx context.Context, aggregateID string) ([]DomainEvent, error)
}
