package v1

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/turao/go-ddd/events"
)

type IntegrationEvent struct {
	baseEvent     *BaseEvent
	correlationID string
}

var _ events.IntegrationEvent = (*IntegrationEvent)(nil)

var (
	ErrInvalidCorrelationID = errors.New("invalid correlation id")
)

func NewIntegrationEvent(event *BaseEvent, correlationID string) (*IntegrationEvent, error) {
	if correlationID == "" {
		return nil, ErrInvalidCorrelationID
	}

	return &IntegrationEvent{
		baseEvent:     event,
		correlationID: correlationID,
	}, nil
}

func (e IntegrationEvent) ID() string {
	return e.baseEvent.id
}

func (e IntegrationEvent) Name() string {
	return e.baseEvent.name
}

func (e IntegrationEvent) OccurredAt() time.Time {
	return e.baseEvent.occuredAt
}

func (e IntegrationEvent) CorrelationID() string {
	return e.correlationID
}

func (e IntegrationEvent) MarshalJSON() ([]byte, error) {
	payload := struct {
		BaseEvent     *BaseEvent `json:"baseEvent"`
		CorrelationID string     `json:"correlationId"`
	}{
		BaseEvent:     e.baseEvent,
		CorrelationID: e.correlationID,
	}
	return json.Marshal(payload)
}

func (e *IntegrationEvent) UnmarshalJson(data []byte) error {
	payload := struct {
		BaseEvent     *BaseEvent `json:"baseEvent"`
		CorrelationID string     `json:"correlationId"`
	}{}

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	e.baseEvent = payload.BaseEvent
	e.correlationID = payload.CorrelationID
	return nil
}
