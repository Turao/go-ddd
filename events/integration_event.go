package events

import "errors"

type integrationEvent struct {
	Event
	correlationID string
}

var _ IntegrationEvent = (*integrationEvent)(nil)

var (
	ErrInvalidCorrelationID = errors.New("invalid correlation id")
)

func NewIntegrationEvent(event Event, correlationID string) (*integrationEvent, error) {
	if correlationID == "" {
		return nil, ErrInvalidCorrelationID
	}

	return &integrationEvent{
		Event:         event,
		correlationID: correlationID,
	}, nil
}

func (ie integrationEvent) CorrelationID() string {
	return ie.correlationID
}
