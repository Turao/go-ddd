package api

import (
	"time"
)

type BaseEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	OccurredAt time.Time `json:"occurredAt"`
}

type DomainEvent struct {
	BaseEvent
	AggregateID string `json:"aggregateId"`
}

type IntegrationEvent struct {
	DomainEvent
	CorrelationID string `json:"correlationId"`
}
