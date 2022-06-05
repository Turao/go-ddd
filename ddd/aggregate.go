package ddd

import (
	"context"
)

type Aggregate interface {
	ID() string
	HandleEvent(ctx context.Context, evt DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) ([]DomainEvent, error)
	MarshalJSON() ([]byte, error)    // aggregate must be serializable due to snapshots
	UnmarshalJSON(data []byte) error // aggregate must be deserializable due to snapshots
}
