package events

type EventStore interface {
	Push(event Event) error
	Events() []Event
	FilterByAggregateID(aggregateID string) []Event
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

func (ims *inMemoryStore) Push(event Event) error {
	ims.evts = append(ims.evts, event)
	return nil
}

func (ims *inMemoryStore) Events() []Event {
	return ims.evts
}

func (ims *inMemoryStore) FilterByAggregateID(aggregateID string) []Event {
	var filtered []Event
	for _, evt := range ims.evts {
		if evt.(DomainEvent).AggregateID() == aggregateID {
			filtered = append(filtered, evt)
		}
	}

	return filtered
}
