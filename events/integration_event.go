package events

import (
	"encoding/json"
	"errors"
	"time"
)

type integrationEvent struct {
	domainEvent   *domainEvent
	correlationID string
}

var (
	ErrInvalidCorrelationID = errors.New("invalid correlation id")
)

func NewIntegrationEvent(name string, aggregateId string, correlationID string) (*integrationEvent, error) {
	de, err := NewDomainEvent(name, aggregateId)
	if err != nil {
		return nil, err
	}

	if correlationID == "" {
		return nil, ErrInvalidCorrelationID
	}

	return &integrationEvent{
		domainEvent:   de,
		correlationID: correlationID,
	}, nil
}

func (ie integrationEvent) ID() string {
	return ie.domainEvent.ID()
}

func (ie integrationEvent) Name() string {
	return ie.domainEvent.Name()
}

func (ie integrationEvent) OccuredAt() time.Time {
	return ie.domainEvent.OccuredAt()
}

func (ie integrationEvent) AggregateID() string {
	return ie.domainEvent.AggregateID()
}

func (ie integrationEvent) CorrelationID() string {
	return ie.correlationID
}

func (ie integrationEvent) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		OccuredAt     time.Time `json:"occurredAt"`
		CorrelationID string    `json:"correlationID"`
	}{
		ID:            ie.domainEvent.ID(),
		Name:          ie.domainEvent.Name(),
		OccuredAt:     ie.domainEvent.OccuredAt(),
		CorrelationID: ie.correlationID,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
