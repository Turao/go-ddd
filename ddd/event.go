package ddd

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/turao/go-ddd/events"
)

type DomainEvent interface {
	events.Event
	AggregateID() string
}

type domainEvent struct {
	events.Event
	aggregateID string
}

var _ DomainEvent = (*domainEvent)(nil)

var (
	ErrInvalidAggregateID = errors.New("invalid aggregate id")
)

func NewDomainEvent(event events.Event, aggregateID string) (*domainEvent, error) {
	if aggregateID == "" {
		return nil, ErrInvalidAggregateID
	}

	return &domainEvent{
		Event:       event,
		aggregateID: aggregateID,
	}, nil
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
		ID:          de.ID(),
		Name:        de.Name(),
		OccuredAt:   de.OccurredAt(),
		AggregateID: de.aggregateID,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
