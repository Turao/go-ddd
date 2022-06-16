package inmemory

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/turao/go-ddd/ddd"
)

type store struct {
	evts map[string][]ddd.DomainEvent
}

var _ ddd.DomainEventStore = (*store)(nil)

var (
	ErrExpectedVersionNotSatisfied = errors.New("expected version does not match event store state")
)

func NewStore() (*store, error) {
	return &store{
		evts: make(map[string][]ddd.DomainEvent, 0),
	}, nil
}

func (ims *store) Push(ctx context.Context, evt ddd.DomainEvent) error {
	// get aggregate events
	aggregateID := evt.AggregateID()
	evts, found := ims.evts[aggregateID]
	if !found {
		evts = make([]ddd.DomainEvent, 0)
		ims.evts[aggregateID] = evts
	}

	// push event into event stream
	evts = append(evts, evt)
	ims.evts[aggregateID] = evts
	return nil
}

func (ims store) Events(ctx context.Context, aggregateID string) ([]ddd.DomainEvent, error) {
	evts, found := ims.evts[aggregateID]
	if !found {
		evts = make([]ddd.DomainEvent, 0)
		ims.evts[aggregateID] = evts
	}

	return evts, nil
}

func (ims store) MarshalJSON() ([]byte, error) {
	events := struct {
		Events map[string][]ddd.DomainEvent `json:"events"`
	}{
		Events: ims.evts,
	}

	return json.MarshalIndent(events, "", " ")
}
