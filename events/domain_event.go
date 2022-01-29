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

var (
	ErrInvalidAggregateID = errors.New("invalid aggregate id")
)

func NewDomainEvent(name string, aggregateID string) (*domainEvent, error) {
	baseEvent, err := NewBaseEvent(name)
	if err != nil {
		return nil, err
	}

	if aggregateID == "" {
		return nil, ErrInvalidAggregateID
	}

	return &domainEvent{
		baseEvent:   baseEvent,
		aggregateID: aggregateID,
	}, nil
}

func (de domainEvent) ID() string {
	return de.baseEvent.ID()
}

func (de domainEvent) Name() string {
	return de.baseEvent.Name()
}

func (de domainEvent) OccuredAt() time.Time {
	return de.baseEvent.OccuredAt()
}

func (de domainEvent) AggregateID() string {
	return de.aggregateID
}

func (de domainEvent) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		OccuredAt   time.Time `json:"occurredAt"`
		AggregateID string    `json:"aggregateId"`
	}{
		ID:          de.baseEvent.ID(),
		Name:        de.baseEvent.Name(),
		OccuredAt:   de.baseEvent.OccuredAt(),
		AggregateID: de.aggregateID,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
