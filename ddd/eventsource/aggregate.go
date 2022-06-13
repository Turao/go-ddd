package eventsource

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/turao/go-ddd/ddd"
)

type aggregate struct {
	aggregate ddd.Aggregate
	version   int

	uncommittedEvents []ddd.DomainEvent
	EventStore        ddd.DomainEventStore
}

var _ ddd.Aggregate = (*aggregate)(nil)

func NewAggregate(agg ddd.Aggregate, es ddd.DomainEventStore) (*aggregate, error) {
	root := &aggregate{
		aggregate:         agg,
		version:           0,
		uncommittedEvents: make([]ddd.DomainEvent, 0),
		EventStore:        es,
	}

	return root, nil
}

func (root aggregate) ID() string {
	return root.aggregate.ID()
}

func (root aggregate) Version() int {
	return root.version
}

func (root *aggregate) HandleEvent(ctx context.Context, evt ddd.DomainEvent) error {
	err := root.aggregate.HandleEvent(ctx, evt)
	if err != nil {
		return err
	}
	root.version += 1
	return nil
}

func (root *aggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	evts, err := root.aggregate.HandleCommand(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// store events in memory until they get committed
	root.uncommittedEvents = evts
	return evts, nil
}

// ReplayEvents fetches all events from the EventStore and executes them in order
func (root *aggregate) ReplayEvents() error {
	// limit how long to wait for the re-creation of the aggregate root
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("fetching all events")
	evts, err := root.EventStore.Events(ctx, root.aggregate.ID())
	if err != nil {
		return err
	}

	log.Println("replaying aggregate events")
	for _, evt := range evts {
		err := root.aggregate.HandleEvent(ctx, evt.(ddd.DomainEvent)) // todo: can we cast these events?
		if err != nil {
			return err
		}
	}

	return nil
}

// CommitEvents flushes all events within the aggregate root's event store
func (root *aggregate) CommitEvents() error {
	evts := root.uncommittedEvents

	// aggregate root version should increment for each event that gets committed
	root.version += len(evts)

	// limit how long to wait for the re-creation of the aggregate root
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, evt := range evts {
		err := root.EventStore.Push(ctx, evt, root.version)
		if err != nil {
			return err
		}
	}

	return nil
}

func (root *aggregate) MarshalJSON() ([]byte, error) {
	snapshot := struct {
		ID        string        `json:"id"`
		Version   int           `json:"version"`
		Aggregate ddd.Aggregate `json:"aggregate"`
	}{
		ID:        root.ID(),
		Version:   root.Version(),
		Aggregate: root.aggregate,
	}

	return json.Marshal(snapshot)
}

func (root *aggregate) UnmarshalJSON(data []byte) error {
	var snapshot struct {
		ID      string `json:"id"`
		Version int    `json:"version"`
		// we cannot unmarshal 'aggregate' into an interface (IDK why yet)
		// so we have to delay aggregate unmarshalling by retrieving the raw json object instead
		Aggregate json.RawMessage `json:"aggregate"`
	}
	err := json.Unmarshal(data, &snapshot)
	if err != nil {
		return err
	}

	root.version = snapshot.Version
	// here we call the concrete implementation
	// which in turn unmarshals the raw message into the correct type
	err = root.aggregate.UnmarshalJSON(snapshot.Aggregate)
	if err != nil {
		return err
	}

	return nil
}
