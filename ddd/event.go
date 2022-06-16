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
	AggregateName() string
}

type domainEvent struct {
	events.Event
	aggregateID   string
	aggregateName string
}

var _ DomainEvent = (*domainEvent)(nil)

var (
	ErrInvalidAggregateID   = errors.New("invalid aggregate id")
	ErrInvalidAggregateName = errors.New("invalid aggregate name")
)

func NewDomainEvent(event events.Event, aggregateID string, aggregateName string) (*domainEvent, error) {
	if aggregateID == "" {
		return nil, ErrInvalidAggregateID
	}

	if aggregateName == "" {
		return nil, ErrInvalidAggregateID
	}

	return &domainEvent{
		Event:         event,
		aggregateID:   aggregateID,
		aggregateName: aggregateName,
	}, nil
}

func (de domainEvent) AggregateID() string {
	return de.aggregateID
}

func (de domainEvent) AggregateName() string {
	return de.aggregateName
}

func (de domainEvent) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		OccuredAt     time.Time `json:"occurredAt"`
		AggregateID   string    `json:"aggregateId"`
		AggregateName string    `json:"aggregateName"`
	}{
		ID:            de.ID(),
		Name:          de.Name(),
		OccuredAt:     de.OccurredAt(),
		AggregateID:   de.aggregateID,
		AggregateName: de.aggregateName,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
