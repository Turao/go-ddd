package events

import (
	"time"
)

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccurredAt() time.Time
}

type IntegrationEvent interface {
	Event
	CorrelationID() string
}
