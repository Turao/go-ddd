package inmemory

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/ddd"
)

type inMemoryStore struct {
	evts []ddd.DomainEvent
}

var _ ddd.DomainEventStore = (*inMemoryStore)(nil)

var (
	ErrExpectedVersionNotSatisfied = errors.New("expected version does not match event store state")
)

func NewInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		evts: make([]ddd.DomainEvent, 0),
	}, nil
}

func (ims *inMemoryStore) Push(ctx context.Context, evt ddd.DomainEvent, expectedVersion int) error {
	version := len(ims.evts) + 1
	if version != expectedVersion {
		return ErrExpectedVersionNotSatisfied
	}

	ims.evts = append(ims.evts, evt)
	return nil
}

func (ims inMemoryStore) Events(ctx context.Context) ([]ddd.DomainEvent, error) {
	return ims.evts, nil
}
