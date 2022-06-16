package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
	v1 "github.com/turao/go-ddd/events/v1"
)

type UserRegisteredEventPublisher interface {
	Publish(ctx context.Context, event UserRegisteredEvent) error
}

type UserRegisteredEvent struct {
	*v1.IntegrationEvent

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
