package events

import (
	"encoding/json"
	"errors"
	"time"
)

type domainEvent struct {
	baseEvent   *baseEvent
	aggregateID string
}

func NewDomainEvent(name string, aggregateID string) (*domainEvent, error) {
	baseEvent, err := NewBaseEvent(name)
	if err != nil {
		return nil, err
	}

	if aggregateID == "" {
		return nil, errors.New("aggregate ID must not be empty")
	}

	return &domainEvent{
		baseEvent:   baseEvent,
		aggregateID: aggregateID,
	}, nil
}

func (de *domainEvent) ID() string {
	return de.baseEvent.id
}

func (de *domainEvent) Name() string {
	return de.baseEvent.name
}

func (de *domainEvent) OccuredAt() time.Time {
	return de.baseEvent.occuredAt
}

func (de *domainEvent) AggregateID() string {
	return de.aggregateID
}

func (de *domainEvent) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		OccuredAt   time.Time `json:"occurredAt"`
		AggregateID string    `json:"aggregateId"`
	}{
		ID:          de.baseEvent.id,
		Name:        de.baseEvent.name,
		OccuredAt:   de.baseEvent.occuredAt,
		AggregateID: de.aggregateID,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
