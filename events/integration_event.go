package events

import "errors"

type IntegrationEvent struct {
	Event
	correlationID string
}

// var _ IntegrationEvent = (*IntegrationEvent)(nil)

var (
	ErrInvalidCorrelationID = errors.New("invalid correlation id")
)

func NewIntegrationEvent(event Event, correlationID string) (*IntegrationEvent, error) {
	if correlationID == "" {
		return nil, ErrInvalidCorrelationID
	}

	return &IntegrationEvent{
		Event:         event,
		correlationID: correlationID,
	}, nil
}

func (ie IntegrationEvent) CorrelationID() string {
	return ie.correlationID
}
