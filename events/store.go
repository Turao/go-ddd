package events

import "context"

type EventStore interface {
	Push(ctx context.Context, evt Event) error
	Events(ctx context.Context) ([]Event, error)
}

// -- in memory implementation --

type inMemoryStore struct {
	evts []Event
}

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		evts: make([]Event, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, evt Event) error {
	ims.evts = append(ims.evts, evt)
	return nil
}

func (ims inMemoryStore) Events(ctx context.Context) ([]Event, error) {
	return ims.evts, nil
}
