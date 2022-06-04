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
	HandleEvent(ctx context.Context, evt DomainEvent) error
	HandleCommand(ctx context.Context, cmd interface{}) error
	CommitEvents() error
}

type root struct {
	aggregate  Aggregate
	version    int
	EventStore events.EventStore
}

var _ AggregateRoot = (*root)(nil)

func NewAggregateRoot(agg Aggregate, es events.EventStore) (*root, error) {
	// limit how long to wait for the re-creation of the aggregate root
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	root := &root{
		aggregate:  agg,
		version:    0,
		EventStore: es,
	}

	// fetch and apply all events from event store
	log.Println("fetching events")
	evts, err := es.Events(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("replaying events")
	for _, evt := range evts {
		err := agg.HandleEvent(ctx, evt.(DomainEvent)) // todo: can we cast these events?
		if err != nil {
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
	log.Printf("handling event - %s", evt.Name())
	defer log.Printf("event handled - %s", evt.Name())
	err := root.aggregate.HandleEvent(ctx, evt)
	if err != nil {
		return err
	}
	root.version += 1
	return nil
}

func (root *root) HandleCommand(ctx context.Context, cmd interface{}) error {
	log.Println("handling command")
	defer log.Println("command handled")
	evts, err := root.aggregate.HandleCommand(ctx, cmd)
	if err != nil {
		return err
	}

	// aggregate root version should increment for each event generated
	// to be consistent with the event handler behavior
	root.version += len(evts)

	for _, evt := range evts {
		err = root.EventStore.Push(ctx, evt, root.version)
		if err != nil {
			return err
		}
	}
	return nil
}

func (root *root) CommitEvents() error {
	log.Println("commiting events - todo!") // todo
	return nil
}
