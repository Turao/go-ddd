package ddd

import (
	"context"
	"encoding/json"
)

type Aggregate interface {
	ID() string

	HandleEvent(ctx context.Context, evt DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) ([]DomainEvent, error)

	json.Marshaler
	json.Unmarshaler
}
