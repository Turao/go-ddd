package events

import "time"

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccuredAt() time.Time
}

// ----

type AggregateID = string

type DomainEvent interface {
	Event
	AggregateID() AggregateID
}

// ----

type CorrelationID = string

type IntegrationEvent interface {
	Event
	CorrelationID() CorrelationID
}
