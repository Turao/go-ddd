package api

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/turao/go-ddd/events"
	v1 "github.com/turao/go-ddd/events/v1"
)

type UserRegisteredEventPublisher interface {
	Publish(ctx context.Context, event UserRegisteredEvent) error
}

type UserRegisteredEvent struct {
	IntegrationEvent *v1.IntegrationEvent `json:"integrationEvent"`

	UserID string `json:"userId"`
}

var _ events.IntegrationEvent = (*UserRegisteredEvent)(nil)

const UserRegisteredEventName = "user.registered"

func NewUserRegisteredEvent(correlationID string, userID string) (*UserRegisteredEvent, error) {
	event, err := v1.NewEvent(UserRegisteredEventName)
	if err != nil {
		return nil, err
	}

	ie, err := v1.NewIntegrationEvent(event, correlationID)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, errors.New("invalid user id")
	}

	return &UserRegisteredEvent{
		IntegrationEvent: ie,
		UserID:           userID,
	}, nil
}

func (e UserRegisteredEvent) ID() string {
	return e.IntegrationEvent.ID()
}

func (e UserRegisteredEvent) Name() string {
	return e.IntegrationEvent.Name()
}

func (e UserRegisteredEvent) CorrelationID() string {
	return e.IntegrationEvent.CorrelationID()
}

func (e UserRegisteredEvent) OccurredAt() time.Time {
	return e.IntegrationEvent.OccurredAt()
}

func (e UserRegisteredEvent) MarshalJSON() ([]byte, error) {
	payload := struct {
		IntegrationEvent *v1.IntegrationEvent `json:"integrationEvent"`
		UserID           string               `json:"userId"`
	}{
		IntegrationEvent: e.IntegrationEvent,
		UserID:           e.UserID,
	}

	return json.Marshal(payload)
}

func (e *UserRegisteredEvent) UnmarshalJSON(data []byte) error {
	payload := struct {
		IntegrationEvent *v1.IntegrationEvent `json:"integrationEvent"`
		UserID           string               `json:"userId"`
	}{}

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	e.IntegrationEvent = payload.IntegrationEvent
	e.UserID = payload.UserID
	return nil
}
