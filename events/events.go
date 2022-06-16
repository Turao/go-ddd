package events

import (
	"time"
)

type Event interface {
	ID() string
	Name() string
	OccurredAt() time.Time
}

type IntegrationEvent interface {
	Event
	CorrelationID() string
}
