package ddd

import (
	"context"
	"log"
	"time"

	"github.com/turao/go-ddd/events"
)

type AggregateRoot interface {
	ID() string
	Version() int

	HandleCommand(ctx context.Context, cmd interface{}) error

	HandleEvent(ctx context.Context, evt DomainEvent) error
	ReplayEvents() error
	CommitEvents() error

	TakeSnapshot() ([]byte, error)
	FromSnapshot(data []byte) error
}

type root struct {
	aggregate  Aggregate
	version    int
	EventStore events.EventStore
}

var _ AggregateRoot = (*root)(nil)

type AggregateRootOption = func(root *root) error

func WithEventStore(es events.EventStore) AggregateRootOption {
	return func(root *root) error {
		root.EventStore = es
		return nil
	}
}

func NewAggregateRoot(agg Aggregate, opts ...AggregateRootOption) (*root, error) {
	root := &root{
		aggregate:  agg,
		version:    0,
		EventStore: nil,
	}

	for _, opt := range opts {
		if err := opt(root); err != nil {
			return nil, err
		}
	}

	return root, nil
}

func (root root) ID() string {
	return root.aggregate.ID() // !risk of null pointers
}

func (root root) Version() int {
	return root.version
}

func (root *root) HandleEvent(ctx context.Context, evt DomainEvent) error {
	err := root.aggregate.HandleEvent(ctx, evt)
	if err != nil {
		return err
	}
	root.version += 1
	return nil
}

func (root *root) HandleCommand(ctx context.Context, cmd interface{}) error {
	evts, err := root.aggregate.HandleCommand(ctx, cmd)
	if err != nil {
		return err
	}

	// aggregate root version should increment for each event generated
	// to be consistent with the event handler behavior
	root.version += len(evts)

	if root.EventStore == nil {
		return nil
	}

	// push events into store
	for _, evt := range evts {
		err = root.EventStore.Push(ctx, evt, root.version)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReplayEvents fetches all events from the EventStore and executes them in order
func (root *root) ReplayEvents() error {
	// limit how long to wait for the re-creation of the aggregate root
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if root.EventStore == nil {
		return nil
	}

	log.Println("fetching events")
	evts, err := root.EventStore.Events(ctx)
	if err != nil {
		return err
	}

	log.Println("replaying events")
	for _, evt := range evts {
		err := root.aggregate.HandleEvent(ctx, evt.(DomainEvent)) // todo: can we cast these events?
		if err != nil {
			return err
		}
	}

	return nil
}

// CommitEvents flushes all events within the aggregate root's event store
func (root *root) CommitEvents() error {
	if root.EventStore == nil {
		return nil
	}

	log.Panic("method not implemented") // todo
	return nil
}

// TakeSnapshot serializes the underlying aggregate
func (root *root) TakeSnapshot() ([]byte, error) {
	return root.aggregate.MarshalJSON()
}

// FromSnapshot deserializes the underlying aggregate
func (root *root) FromSnapshot(data []byte) error {
	return root.aggregate.UnmarshalJSON(data)
}
