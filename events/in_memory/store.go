package in_memory

import (
	"context"

	"github.com/turao/go-ddd/events"
)

// -- in memory implementation --

type inMemoryStore struct {
	evts []events.Event
}

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		evts: make([]events.Event, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, evt events.Event) error {
	ims.evts = append(ims.evts, evt)
	return nil
}

func (ims inMemoryStore) Events(ctx context.Context) ([]events.Event, error) {
	return ims.evts, nil
}
