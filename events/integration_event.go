package events

import (
	"encoding/json"
	"errors"
	"time"
)

type integrationEvent struct {
	baseEvent     *baseEvent
	correlationID string
}

func NewIntegrationEvent(name string, correlationID string) (*integrationEvent, error) {
	baseEvent, err := NewBaseEvent(name)
	if err != nil {
		return nil, err
	}

	if correlationID == "" {
		return nil, errors.New("correlation ID must not be empty")
	}

	return &integrationEvent{
		baseEvent:     baseEvent,
		correlationID: correlationID,
	}, nil
}

func (ie integrationEvent) ID() string {
	return ie.baseEvent.id
}

func (ie integrationEvent) Name() string {
	return ie.baseEvent.name
}

func (ie integrationEvent) OccuredAt() time.Time {
	return ie.baseEvent.occuredAt
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
		ID:            ie.baseEvent.id,
		Name:          ie.baseEvent.name,
		OccuredAt:     ie.baseEvent.occuredAt,
		CorrelationID: ie.correlationID,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
