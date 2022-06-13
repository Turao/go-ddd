package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
)

type UserRegisteredEventPublisher interface {
	Publish(ctx context.Context, event UserRegisteredEvent) error
}

type UserRegisteredEvent struct {
	*events.IntegrationEvent

	UserID string `json:"userId"`
}

// var _ events.IntegrationEvent = (*UserRegisteredEvent)(nil)

const UserRegisteredEventName = "user.registered"

func NewUserRegisteredEvent(correlationID string, userID string) (*UserRegisteredEvent, error) {
	event, err := events.NewEvent(UserRegisteredEventName)
	if err != nil {
		return nil, err
	}

	ie, err := events.NewIntegrationEvent(event, correlationID)
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
