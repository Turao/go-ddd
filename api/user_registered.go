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
	IntegrationEvent
	UserID string `json:"userId"`
}

const (
	UserRegisteredEventName = "user.registered"
)

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUserRegisteredEvent(correlationID string, userID string) (*UserRegisteredEvent, error) {

	ie, err := events.NewIntegrationEvent(UserRegisteredEventName, userID, correlationID)
	if err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &UserRegisteredEvent{
		IntegrationEvent: IntegrationEvent{
			DomainEvent: DomainEvent{
				BaseEvent: BaseEvent{
					ID:         ie.ID(),
					Name:       ie.Name(),
					OccurredAt: ie.OccuredAt(),
				},
				AggregateID: ie.AggregateID(),
			},
			CorrelationID: ie.CorrelationID(),
		},
		UserID: userID,
	}, nil
}
