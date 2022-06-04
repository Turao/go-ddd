package ddd

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/events/in_memory"
)

type root struct {
	id            string
	version       int
	eventStore    events.EventStore
	handleEvent   EventHandlerFunc
	handleCommand CommandHandlerFunc
}

func NewAggregateRoot() (*root, error) {
	log.Println("creating aggregate root")
	store, err := in_memory.NewInMemoryStore()
	if err != nil {
		return nil, err
	}

	root := &root{
		id:         uuid.NewString(),
		version:    0,
		eventStore: store,
	}

	return root, nil
}

func FromSnapshot(snap Snapshot) (*root, error) {
	log.Println("creating aggregate root from snapshot: ", snap)
	store, err := in_memory.NewInMemoryStore()
	if err != nil {
		return nil, err
	}

	root := &root{
		id:         snap.ID,
		version:    snap.Version,
		eventStore: store,
	}

	return root, nil
}

func (root *root) ReplayEvents(ctx context.Context) error {
	// fetch and apply all events from event store
	log.Println("fetching events from event store")
	evts, err := root.eventStore.Events(ctx)
	if err != nil {
		return err
	}

	log.Println("replaying events")
	for _, evt := range evts {
		err := root.handleEvent(ctx, evt.(DomainEvent)) // todo: can we cast these events?
		if err != nil {
			return err
		}
	}

	return nil
}

type EventHandlerFunc = func(ctx context.Context, event DomainEvent) error

func (root *root) RegisterEventHandlerFunc(handler EventHandlerFunc) {
	root.handleEvent = handler
}

func (root *root) HandleEvent(ctx context.Context, evt DomainEvent) error {
	log.Printf("handling event - %s", evt.Name())
	defer log.Printf("event handled - %s", evt.Name())

	if err := root.handleEvent(ctx, evt); err != nil {
		return err
	}

	root.version += 1
	return nil
}

type CommandHandlerFunc = func(ctx context.Context, cmd interface{}) ([]DomainEvent, error)

func (root *root) RegisterCommandHandlerFunc(handler CommandHandlerFunc) {
	root.handleCommand = handler
}

func (root *root) HandleCommand(ctx context.Context, cmd interface{}) error {
	log.Println("handling command")
	defer log.Println("command handled")

	evts, err := root.handleCommand(ctx, cmd)
	if err != nil {
		return err
	}

	// aggregate root version should increment for each event generated
	// to be consistent with the event handler behavior
	root.version += len(evts)

	for _, evt := range evts {
		err = root.eventStore.Push(ctx, evt, root.version)
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
