package api

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type UserRegisteredEventPublisher interface {
	Publish(ctx context.Context, event UserRegisteredEvent) error
}

type UserRegisteredEvent struct {
	events.IntegrationEvent
	UserID string `json:"userId"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUserRegisteredEvent(userID string) (*UserRegisteredEvent, error) {
	ie, err := events.NewIntegrationEvent("user.registered", uuid.NewString())
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &UserRegisteredEvent{
		IntegrationEvent: ie,
		UserID:           userID,
	}, nil
}
