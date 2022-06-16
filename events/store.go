package events

import "context"

type EventStore interface {
	Push(ctx context.Context, evt Event) error
	Events(ctx context.Context) ([]Event, error)
}
