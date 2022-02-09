package in_memory

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
)

// -- in memory implementation --

type inMemoryStore struct {
	evts []events.Event
}

var _ events.EventStore = (*inMemoryStore)(nil)

var (
	ErrExpectedVersionNotSatisfied = errors.New("expected version does not match event store state")
)

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		evts: make([]events.Event, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, evt events.Event, expectedVersion int) error {
	version := len(ims.evts) + 1
	if version != expectedVersion {
		return ErrExpectedVersionNotSatisfied
	}

	ims.evts = append(ims.evts, evt)
	return nil
}

func (ims inMemoryStore) Events(ctx context.Context) ([]events.Event, error) {
	return ims.evts, nil
}

func (ims inMemoryStore) EventsByAggregateID(ctx context.Context, aggregateID events.AggregateID) ([]events.DomainEvent, error) {
	es := make([]events.DomainEvent, 0)
	for _, e := range ims.evts {
		de, ok := e.(events.DomainEvent)
		if ok {
			if de.AggregateID() == aggregateID {
				es = append(es, de)
			}
		}
	}
	return es, nil
}
