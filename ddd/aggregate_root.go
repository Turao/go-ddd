package ddd

import (
	"context"
	"log"
	"time"

	"github.com/turao/go-ddd/events"
)

type AggregateRoot struct {
	Aggregate
	version    int
	EventStore events.EventStore
}

func NewAggregateRoot(a Aggregate, es events.EventStore) (*AggregateRoot, error) {
	// limit how long to wait for the re-creation of the aggregate root
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	root := &AggregateRoot{
		Aggregate:  a,
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
		err := a.HandleEvent(ctx, evt.(events.DomainEvent)) // todo: can we cast these events?
		if err != nil {
			return nil, err
		}
	}

	return root, nil
}

func (a *AggregateRoot) HandleEvent(ctx context.Context, evt events.DomainEvent) error {
	log.Printf("handling event - %s", evt.Name())
	defer log.Printf("event handled - %s", evt.Name())
	err := a.Aggregate.HandleEvent(ctx, evt)
	if err != nil {
		return err
	}
	a.version += 1
	return nil
}

func (a *AggregateRoot) HandleCommand(ctx context.Context, cmd interface{}) error {
	log.Println("handling command")
	defer log.Println("command handled")
	evts, err := a.Aggregate.HandleCommand(ctx, cmd)
	if err != nil {
		return err
	}

	// aggregate root version should increment for each command successfully handled
	a.version += len(evts)

	for _, evt := range evts {
		err = a.EventStore.Push(ctx, evt, a.version)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AggregateRoot) CommitEvents() error {
	log.Println("commiting events - todo!") // todo
	return nil
}
