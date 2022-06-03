package ddd

import (
	"context"
)

type Aggregate interface {
	ID() string // aggregate ID must be serializable
	HandleEvent(ctx context.Context, evt DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) ([]DomainEvent, error)
}
