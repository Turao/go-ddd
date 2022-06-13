package inmemory

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/turao/go-ddd/ddd"
)

type inMemoryStore struct {
	evts map[string][]ddd.DomainEvent
}

var _ ddd.DomainEventStore = (*inMemoryStore)(nil)

var (
	ErrExpectedVersionNotSatisfied = errors.New("expected version does not match event store state")
)

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		evts: make(map[string][]ddd.DomainEvent, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, evt ddd.DomainEvent, expectedVersion int) error {
	// get aggregate events
	aggregateID := evt.AggregateID()
	evts, found := ims.evts[aggregateID]
	if !found {
		evts = make([]ddd.DomainEvent, 0)
		ims.evts[aggregateID] = evts
	}

	// validate expected version
	version := len(evts) + 1
	if version != expectedVersion {
		return ErrExpectedVersionNotSatisfied
	}

	// push event into event stream
	evts = append(evts, evt)
	ims.evts[aggregateID] = evts
	return nil
}

func (ims inMemoryStore) Events(ctx context.Context, aggregateID string) ([]ddd.DomainEvent, error) {
	evts, found := ims.evts[aggregateID]
	if !found {
		evts = make([]ddd.DomainEvent, 0)
		ims.evts[aggregateID] = evts
	}

	return evts, nil
}

func (ims inMemoryStore) MarshalJSON() ([]byte, error) {
	events := struct {
		Events map[string][]ddd.DomainEvent `json:"events"`
	}{
		Events: ims.evts,
	}

	return json.MarshalIndent(events, "", " ")
}
