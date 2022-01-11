package events

import (
	"errors"
	"time"
)

type DomainEvent interface {
	AggregateID() string
}

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

func (e *domainEvent) ID() EventID {
	return e.baseEvent.id
}

func (e *domainEvent) Name() string {
	return e.baseEvent.name
}

func (e *domainEvent) OccuredAt() time.Time {
	return e.baseEvent.occuredAt
}

func (e *domainEvent) AggregateID() string {
	return e.aggregateID
}
