package aggregate

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type Aggregate interface {
	ID() string // aggregate ID must be serializable
	HandleEvent(ctx context.Context, evt events.DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) ([]events.DomainEvent, error)
}
