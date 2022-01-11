package events

import (
	"context"
)

type EventStore interface {
	Push(ctx context.Context, event Event) error
	Take(ctx context.Context, events int) []Event
}

// -- in memory implementation --

type inMemoryStore struct {
	events []Event
}

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		events: make([]Event, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, event Event) error {
	ims.events = append(ims.events, event)
	return nil
}

func (ims *inMemoryStore) Take(ctx context.Context, events int) []Event {
	toBeTaken := events
	if toBeTaken > len(ims.events) {
		toBeTaken = len(ims.events) // take as much as possible
	}

	taken := ims.events[0:toBeTaken]
	ims.events = ims.events[toBeTaken:len(ims.events)]
	return taken
}
